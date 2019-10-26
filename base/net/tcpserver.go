package net

import (
	"base/log"
	"net"
)

type TcpServer struct {
	listener *net.TCPListener
}

func (this *TcpServer) Bind(address string) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		log.Error.Println("bind fail ", address)
		return err
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Error.Println("listen fail ", address)
		return err
	}
	log.Info.Println("bind success ", address)
	this.listener = listener
	return nil
}

func (this *TcpServer) Accept() (*net.TCPConn, error){
	conn, err := this.listener.AcceptTCP()
	if err != nil {
		log.Error.Println("Accept error")
		return nil, err
	}
	_ = conn.SetKeepAlive(true)
	_ = conn.SetNoDelay(true)
	_ = conn.SetWriteBuffer(64 * 1024)
	_ = conn.SetReadBuffer(64 * 1024)
	return conn, err
}

func (this *TcpServer) Close() error {
	return this.listener.Close()
}