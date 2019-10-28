package net

import (
	"base/log"
	"base/util"
	"bytes"
	"io"
	"net"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

const (
	sendDataMaxSize = 64 * 1024
	cmdDataHeadSize = 4
)

type ITcpTask interface {
	SendData(data []byte)
	SendDataWithHead(head []byte, data []byte)
	ParseMsg(data []byte) bool
	OnClose()
	Close()
}

type TcpTask struct {
	userTcpTask ITcpTask
	isClosed int32
	Conn net.Conn
	sendBuff *ByteBuffer
	sendMutex sync.Mutex
	sendSignal chan bool
}

func NewTcpTask(conn net.Conn) *TcpTask {
	return &TcpTask{
		isClosed: 1,
		Conn: conn,
		sendBuff: NewByteBuffer(),
		sendSignal: make(chan bool, 1),
	}
}

func (this *TcpTask) SetUserTask (userTcpTask ITcpTask) {
	this.userTcpTask = userTcpTask
}

func (this *TcpTask) Start() {
	if !atomic.CompareAndSwapInt32(&this.isClosed, 1, 0) {
		return
	}
	log.Debug.Println("start TcpTask addr: ", this.Conn.RemoteAddr())
	go this.sendLoop()
	go this.recvLoop()
}

func (this *TcpTask) IsClosed() bool {
	return atomic.LoadInt32(&this.isClosed) != 0
}

func (this *TcpTask) Close () {
	if !atomic.CompareAndSwapInt32(&this.isClosed, 0, 1) {
		return
	}
	log.Debug.Println("close TcpTask addr:", this.Conn.RemoteAddr())
	_ = this.Conn.Close()
	this.sendSignal <- true
	close(this.sendSignal)
	this.userTcpTask.OnClose()
}

func (this *TcpTask) SendData (data []byte) {
	if this.IsClosed() {
		log.Error.Println("TcpTask isClosed")
		return
	}
	dataSize := len(data)
	if dataSize == 0 {
		log.Warn.Println("send data size:", dataSize)
		return
	}
	this.sendMutex.Lock()
	curBuffSize := this.sendBuff.RdSize()
	if curBuffSize + dataSize > sendDataMaxSize {
		log.Error.Printf("send buff over limit cur:%d new:%d", curBuffSize, dataSize)
		this.Close()
	}
	this.sendBuff.Append(util.Uint32ToBytes(uint32(dataSize)))
	this.sendBuff.Append(data)
	this.sendMutex.Unlock()
	this.sendSignal <- false
}

func (this *TcpTask) SendDataWithHead(head []byte, data []byte) {
	fullData := new(bytes.Buffer)
	fullData.Write(head)
	fullData.Write(data)
	this.SendData(fullData.Bytes())
}

func (this *TcpTask) recvLoop() {
	defer func() {
		if err := recover(); err != nil {
			log.Error.Println("err:", err, "\n", string(debug.Stack()))
		}
	}()

	defer this.Close()

	var (
		recvBuff  *ByteBuffer = NewByteBuffer()
		totalSize int
		dataSize int
		needNum int
		readNum int
		err error
		msgBuff []byte

	)

	for {
		totalSize = recvBuff.RdSize()
		for totalSize < cmdDataHeadSize {
			needNum = cmdDataHeadSize - totalSize
			// log.Debug.Println("recv head needNum:", needNum)
			readNum, err = this.Conn.Read(recvBuff.WrBuf())
			// readNum, err = io.ReadAtLeast(this.Conn, recvBuff.WrBuf(), needNum)
			if err == io.EOF{
				log.Debug.Println("remote close io eof addr:", this.Conn.RemoteAddr())
				return
			} else if err != nil {
				log.Error.Printf("recv loop addr:%s, err:%T %+v", this.Conn.RemoteAddr(), err, err)
				return
			}
			recvBuff.WrFlip(readNum)
			totalSize = recvBuff.RdSize()
			// log.Debug.Println("recv head readNum:", readNum, " totalSize:", totalSize)
		}
		msgBuff = recvBuff.RdBuf()
		dataSize = int(util.BytesToUint32(msgBuff[:cmdDataHeadSize]))
		if dataSize > sendDataMaxSize {
			log.Error.Println("recv too big data over limit size:", dataSize)
			return
		}
		for totalSize < cmdDataHeadSize + dataSize {
			needNum = cmdDataHeadSize + dataSize - totalSize
			if recvBuff.WrSize() < needNum {
				recvBuff.WrGrow(needNum)
			}
			// log.Debug.Println("recv body needNum:", needNum)
			readNum, err = this.Conn.Read(recvBuff.WrBuf())
			// readNum, err = io.ReadAtLeast(this.Conn, recvBuff.WrBuf(), needNum)
			if err != nil {
				log.Error.Println("recv loop addr:", this.Conn.RemoteAddr(), ", err:", err)
				return
			}
			// log.Debug.Println("recv body readNum:", readNum, " totalSize:", totalSize)
			recvBuff.WrFlip(readNum)
		}
		msgBuff = recvBuff.RdBuf()
		this.userTcpTask.ParseMsg(msgBuff[cmdDataHeadSize:cmdDataHeadSize+dataSize])
		recvBuff.RdFlip(cmdDataHeadSize + dataSize)
	}
}

func (this *TcpTask) sendLoop() {
	defer func() {
		if err := recover(); err != nil {
			log.Error.Println("err: ", err, "\n", string(debug.Stack()))
		}
	}()

	defer this.Close()

	var (
		tmpByte = NewByteBuffer()
		sendNum int
		err error
	)

	for needClose := range this.sendSignal {
		if needClose {
			return
		} else {
			this.sendMutex.Lock()
			if this.sendBuff.RdReady() {
				// 发送数据由发送缓冲区移动到发送协程中
				tmpByte.Append(this.sendBuff.RdBuf()[:this.sendBuff.RdSize()])
				this.sendBuff.Reset()
			}
			this.sendMutex.Unlock()
			if !tmpByte.RdReady() {
				continue
			}
			for tmpByte.RdReady() {
				// 发送完整
				sendNum, err = this.Conn.Write(tmpByte.RdBuf()[:tmpByte.RdSize()])
				if err != nil {
					log.Error.Println("send loop addr:", this.Conn.RemoteAddr(), ", err:", err)
					return
				}
				tmpByte.RdFlip(sendNum)
			}
		}
	}
}