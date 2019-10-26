package usertask

import (
	"base/log"
	"base/net"
	"base/util"
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
	value := util.BytesToUint32(data)
	log.Debug.Println("addr:", this.Conn.RemoteAddr(), " recv data:", data, "value:", value)
	return true
}

func (this *LoginTask) OnClose () {

}