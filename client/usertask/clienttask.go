package usertask

import (
	"base/log"
	"base/net"
	engineNet "net"
)

type Task struct {
	net.TcpTask
}

func NewTask(conn engineNet.Conn) *Task {
	task := &Task{
		TcpTask: *net.NewTcpTask(conn),
	}
	task.SetUserTask(task)
	return task
}

func (this *Task) ParseMsg(data []byte) bool {
	log.Debug.Println("addr: ", this.Conn.RemoteAddr(), " recv data: ", data)
	return true
}

func (this *Task) OnClose() {

}