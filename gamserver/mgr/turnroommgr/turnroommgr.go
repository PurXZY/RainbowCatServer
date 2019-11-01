package turnroommgr

import (
	"gamserver/entity/turnroom"
	"gamserver/i"
	"gamserver/mgr/idmgr"
	"sync"
)

type TurnRoomMgr struct {
	turnRooms map[uint32]*turnroom.TurnRoom
}

var (
	mgr     *TurnRoomMgr
	mgrOnce sync.Once
)

func GetMe() *TurnRoomMgr {
	if mgr == nil {
		mgrOnce.Do(func() {
			mgr = &TurnRoomMgr{
				turnRooms: make(map[uint32]*turnroom.TurnRoom),
			}
		})
	}
	return mgr
}

func (this *TurnRoomMgr) AddNewTurnRoom(owner i.IRoomOwner) {
	uniqId := idmgr.GetMe().GenUniqId()
	room := turnroom.NewTurnRoom(uniqId, owner)
	this.turnRooms[uniqId] = room
	owner.SetTurnRoom(room)
	room.Init()
}
