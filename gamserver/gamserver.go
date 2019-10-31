package gamserver

import (
	"base/log"
	"base/net"
	"gamserver/mgr/datamgr"
	"gamserver/mgr/sessionmgr"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"syscall"
	"time"
)

type GameServer struct {
	server    net.TcpServer
	terminate bool
}

func NewGameServer() *GameServer {
	ser := &GameServer{
		server:    net.TcpServer{},
		terminate: false,
	}
	return ser
}

func (this *GameServer) Start() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Error.Println("err: ", err, "\n", string(debug.Stack()))
		}
	}()
	this.DealWithSignal()
	runtime.GOMAXPROCS(runtime.NumCPU())
	if !this.Init() {
		return false
	}
	for !this.isTerminate() {
		this.MainLoop()
	}
	this.Final()
	return true
}

func (this *GameServer) DealWithSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGPIPE, syscall.SIGHUP)
	go func() {
		for sig := range ch {
			log.Info.Println("receive signal: ", sig)
			switch sig {
			case syscall.SIGPIPE:
				log.Error.Println("SIGPIPE")
			default:
				this.Terminate()
			}
		}
	}()
}

func (this *GameServer) Terminate() {
	this.terminate = true
	_ = this.server.Close()
}

func (this *GameServer) isTerminate() bool {
	return this.terminate
}

func (this *GameServer) Init() bool {
	log.Info.Println("init server")
	rand.Seed(time.Now().Unix())

	var address string = "127.0.0.1:8888"
	err := this.server.Bind(address)
	if err != nil {
		log.Error.Println("bind fail")
		return false
	}
	log.Info.Println("init server success")
	datamgr.GetMe()
	return true
}

func (this *GameServer) MainLoop() {
	conn, err := this.server.Accept()
	if err != nil {
		return
	}
	sessionmgr.GetMe().AddNewSession(conn)
}

func (this *GameServer) Final() {
	log.Info.Println("server final")
}
