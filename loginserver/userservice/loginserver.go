package userservice

import (
	"base/log"
	"base/net"
	"loginserver/usertask"
	"loginserver/usertaskmgr"
	"math/rand"
	"time"
)
type LoginServer struct {
	net.Service
}

func NewLoginServer() *LoginServer{
	ser := &LoginServer{
		*net.NewService(),
	}
	ser.SetUserService(ser)
	return ser
}

func (this *LoginServer) Init() bool {
	log.Info.Println("init server")
	rand.Seed(time.Now().Unix())

	var address string = "127.0.0.1:8888"
	err := this.Server.Bind(address)
	if err != nil {
		log.Error.Println("bind fail")
		return false
	}
	log.Info.Println("init server success")
	return true
}

func (this *LoginServer) MainLoop() {
	conn, err := this.Server.Accept()
	if err != nil {
		return
	}
	loginTask := usertask.NewLoginTask(conn)
	ok := usertaskmgr.GetMe().AddNewPlayerTask(loginTask)
	if ok {
		loginTask.Start()
	}
}

func (this *LoginServer) Final() {
	log.Info.Println("server final")
}