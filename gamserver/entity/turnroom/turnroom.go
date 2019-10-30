package turnroom

import (
	"base/log"
	"base/util"
	"gamserver/i"
	"usercmd"
)

type TurnRoom struct {
	uniqId uint32
	owner  i.ISessionOwner
}

func NewTurnRoom(uniqId uint32, owner i.ISessionOwner) *TurnRoom {
	room := &TurnRoom{
		uniqId: uniqId,
		owner:  owner,
	}
	return room
}

func (this *TurnRoom) Init() {
	log.Info.Printf("new TurnRoom id:%v, owner:%v", this.uniqId, this.owner.GetName())
	msg := usercmd.IntoRoomS2CMsg{
		RoomId: this.uniqId,
	}
	sendData, err := util.EncodeCmd(usercmd.UserCmd_IntoRoomRes, &msg)
	if err != nil {
		return
	}
	this.owner.SendData(sendData)
}
