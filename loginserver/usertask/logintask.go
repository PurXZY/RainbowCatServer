package usertask

import (
	"base/log"
	"base/net"
	engineNet "net"
)

type LoginTask struct {
	net.TcpTask
}

func NewLoginTask(conn engineNet.Conn) *LoginTask {
	task := &LoginTask{
		TcpTask: *net.NewTcpTask(conn),
	}
	task.SetUserTask(task)
	return task
}

func (this *LoginTask) ParseMsg(data []byte) bool {
	log.Debug.Println("addr: ", this.Conn.RemoteAddr(), " recv data: ", data)
	return true
}

func (this *LoginTask) OnClose () {

}