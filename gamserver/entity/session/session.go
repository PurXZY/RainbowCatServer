package session

import (
	"base/log"
	"base/net"
	"base/util"
	"gamserver/i"
	"gamserver/mgr/avatarmgr"
	engineNet "net"
	"usercmd"
)

type Session struct {
	net.TcpTask
	sessionId uint32
	owner     i.ISessionOwner
}

func NewSession(conn engineNet.Conn, sessionId uint32) *Session {
	session := &Session{
		*net.NewTcpTask(conn),
		sessionId,
		nil,
	}
	session.SetSession(session)
	return session
}

func (this *Session) SetOwner(owner i.ISessionOwner) {
	this.owner = owner
}

func (this *Session) ParseMsg(data []byte) bool {
	cmdType := usercmd.UserCmd(util.GetCmdType(data))
	log.Debug.Println("ParseMsg cmdType:", cmdType)
	switch cmdType {
	case usercmd.UserCmd_LoginReq:
		recvCmd, ok := util.DecodeCmd(data, &usercmd.LoginC2SMsg{}).(*usercmd.LoginC2SMsg)
		if !ok {
			log.Error.Println("decode cmd fail cmdType:", cmdType)
			return false
		}
		log.Debug.Println("recv msg UserCmd_LoginReq name:", recvCmd.GetName())
		msg := usercmd.LoginS2CMsg{
			PlayerId: this.sessionId,
		}
		data, err := util.EncodeCmd(usercmd.UserCmd_LoginRes, &msg)
		if err != nil {
			return false
		}
		this.SendData(data)
		avatarmgr.GetMe().AddNewAvatar(this, recvCmd.GetName())
	default:
		this.owner.OnRecvMsg(cmdType, data)
	}
	return true
}

func (this *Session) OnClose() {

}
