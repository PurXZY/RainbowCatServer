package avatar

import (
	"fmt"
	"gamserver/i"
	"gamserver/mgr/turnroommgr"
	"strconv"
)

type Avatar struct {
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
	return avatar
}

func (this *Avatar) ReqIntoRoom() {
	turnroommgr.GetMe().AddNewTurnRoom(this)
}

func (this *Avatar) GetName() string {
	id := strconv.FormatUint(uint64(this.uniqId), 10)
	return fmt.Sprintf("[avatar id:%v name:%v]", id, this.name)
}

func (this *Avatar) SendData(data []byte) {
	this.session.SendData(data)
}