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
	owner i.ISessionOwner
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
	switch cmdType {
	case usercmd.UserCmd_LoginReq:
		recvCmd, ok := util.DecodeCmd(data, &usercmd.LoginC2SMsg{}).(*usercmd.LoginC2SMsg)
		if !ok {
			log.Error.Println("decode cmd fail cmdType:", cmdType)
			return false
		}
		log.Debug.Println("recv msg UserCmd_LoginReq name:", recvCmd.GetName())
		this.sendLoginRes()
		avatarmgr.GetMe().AddNewAvatar(this, recvCmd.GetName())
	case usercmd.UserCmd_IntoRoomReq:
		this.owner.ReqIntoRoom()
	default:
		log.Error.Println("unknown cmdType:", cmdType)
		return false
	}
	return true
}

func (this *Session) OnClose() {

}

func (this *Session) sendLoginRes() {
	msg := usercmd.LoginS2CMsg{
		PlayerId: this.sessionId,
	}
	data, err := util.EncodeCmd(usercmd.UserCmd_LoginRes, &msg)
	if err != nil {
		return
	}
	this.SendData(data)
}
