package usertask

import (
	"base/log"
	"base/net"
	"base/util"
	"loginserver/usertaskmgr"
	engineNet "net"
	"usercmd"
)

type LoginTask struct {
	net.TcpTask
	taskId uint32
}

func NewLoginTask(conn engineNet.Conn) *LoginTask {
	task := &LoginTask{
		TcpTask: *net.NewTcpTask(conn),
	}
	task.SetUserTask(task)
	return task
}

func (this *LoginTask) SetId(id uint32) {
	this.taskId = id
	log.Debug.Printf("add new task id:%v addr:%v", this.taskId, this.Conn.RemoteAddr())

}

func (this *LoginTask) GetId() uint32 {
	return this.taskId
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
		this.sendLoginRes()
	default:
		log.Error.Println("unknown cmdType:", cmdType)
		return false
	}
	return true
}

func (this *LoginTask) OnClose () {
	usertaskmgr.GetMe().DeletePlayerTask(this.taskId)
}

func (this *LoginTask) sendLoginRes() {
	msg := usercmd.LoginS2CMsg{
		PlayerId:this.GetId(),
	}
	data, err := util.EncodeCmd(usercmd.UserCmd_LoginRes, &msg)
	if err != nil {
		return
	}
	this.SendData(data)
}