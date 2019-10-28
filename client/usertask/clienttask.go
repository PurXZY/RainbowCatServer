package usertask

import (
	"base/log"
	"base/net"
	"base/util"
	engineNet "net"
	"usercmd"
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

func (this *Task) SetId(id uint32) {

}

func (this *Task) GetId() uint32 {
	return 0
}


func (this *Task) ParseMsg(data []byte) bool {
	cmdType := usercmd.UserCmd(util.GetCmdType(data))
	switch cmdType {
	case usercmd.UserCmd_LoginRes:
		recvCmd, ok := util.DecodeCmd(data, &usercmd.LoginS2CMsg{}).(*usercmd.LoginS2CMsg)
		if !ok {
			log.Error.Println("decode cmd fail cmdType:", cmdType)
			return false
		}
		log.Debug.Println("recv msg UserCmd_LoginRes id:", recvCmd.GetPlayerId())
	}
	return true
}

func (this *Task) OnClose() {

}