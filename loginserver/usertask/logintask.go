package usertask

import (
	"base/log"
	"base/net"
	"base/util"
	engineNet "net"
	"usercmd"
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
	cmdType := usercmd.UserCmd(util.GetCmdType(data))
	switch cmdType {
	case usercmd.UserCmd_LoginReq:
		recvCmd, ok := util.DecodeCmd(data, &usercmd.LoginC2SMsg{}).(*usercmd.LoginC2SMsg)
		if !ok {
			log.Error.Println("decode cmd fail cmdType:",cmdType)
			return false
		}
		log.Debug.Println("recv msg UserCmd_LoginReq name:", recvCmd.GetName())
	default:
		log.Error.Println("unknown cmdType:", cmdType)
		return false
	}
	return true
}

func (this *LoginTask) OnClose () {

}