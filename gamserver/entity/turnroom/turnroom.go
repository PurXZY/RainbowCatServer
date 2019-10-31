package turnroom

import (
	"base/log"
	"base/util"
	"gamserver/i"
	"usercmd"
)

type TurnRoom struct {
	EntityMgr
	TurnLogic
	uniqId uint32
	owner  i.IRoomOwner
}

func NewTurnRoom(uniqId uint32, owner i.IRoomOwner) *TurnRoom {
	room := &TurnRoom{
		EntityMgr: *NewEntityMgr(),
		uniqId:    uniqId,
		owner:     owner,
	}
	return room
}

func (this *TurnRoom) Init() {
	log.Info.Printf("new TurnRoom id:%v, owner:%v", this.uniqId, this.owner.GetName())

	this.notifyPlayerIntoRoom()
	this.EntityMgr.Init(this)
	this.TurnLogic.Init(this)
}

func (this *TurnRoom) BroadcastMsg(data []byte) {
	this.owner.SendData(data)
}

func (this *TurnRoom) notifyPlayerIntoRoom() {
	msg := usercmd.IntoRoomS2CMsg{
		RoomId: this.uniqId,
	}
	sendData, err := util.EncodeCmd(usercmd.UserCmd_IntoRoomRes, &msg)
	if err != nil {
		return
	}
	this.BroadcastMsg(sendData)
}

