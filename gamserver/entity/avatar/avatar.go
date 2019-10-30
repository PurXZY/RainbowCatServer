package avatar

import (
	"fmt"
	"gamserver/i"
	"strconv"
	"usercmd"
)

type Avatar struct {
	AvatarNetMsgDispatcher
	uniqId  uint32
	session i.ISession
	name    string
}

func NewAvatar(id uint32, session i.ISession, name string) *Avatar {
	avatar := &Avatar{
		uniqId:  id,
		session: session,
		name:    name,
	}
	session.SetOwner(avatar)
	avatar.Init(avatar)
	return avatar
}

func (this *Avatar) OnRecvMsg(cmd usercmd.UserCmd, data []byte) {
	this.dispatcher.Call(cmd, data)
}

func (this *Avatar) GetName() string {
	id := strconv.FormatUint(uint64(this.uniqId), 10)
	return fmt.Sprintf("[avatar id:%v name:%v]", id, this.name)
}

func (this *Avatar) SendData(data []byte) {
	this.session.SendData(data)
}