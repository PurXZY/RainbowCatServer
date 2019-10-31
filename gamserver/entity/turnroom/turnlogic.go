package turnroom

import (
	"base/log"
	"base/util"
	"gamserver/i"
	"sort"
	"usercmd"
)

type TurnLogic struct {
	turnRoom       i.ITurnRoom
	curBigTurn     uint32
	curSmallTurn   uint32
	entityTurnInfo []uint32
}

func (this *TurnLogic) Init(room i.ITurnRoom) {
	this.turnRoom = room
	this.beginFirstTurn()
}

type PosSpeedPair struct {
	pos   uint32
	speed uint32
}
type PairList []PosSpeedPair

func (pair PairList) Swap(i, j int)   { pair[i], pair[j] = pair[j], pair[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].speed < p[j].speed }

func (this *TurnLogic) SortEntitySpeedInfo() {
	speedMap := this.turnRoom.GetAllEntitiesSpeed()
	var plist PairList
	for posIndex, speed := range speedMap {
		plist = append(plist, PosSpeedPair{
			pos:   posIndex,
			speed: speed,
		})
	}
	sort.Sort(plist)
	this.entityTurnInfo = this.entityTurnInfo[:0]
	for _, pair := range plist {
		this.entityTurnInfo = append(this.entityTurnInfo, pair.pos)
	}
}

func (this *TurnLogic) beginFirstTurn() {
	this.curBigTurn = 1
	this.curSmallTurn = 0
	this.SortEntitySpeedInfo()
	this.NotifyCurTurn()
}

func (this *TurnLogic) curTurnEntity() uint32 {
	return this.entityTurnInfo[this.curSmallTurn]
}

func (this TurnLogic) NotifyCurTurn() {
	log.Info.Printf("New Turn Big:%v Small:%v Pos:%v", this.curBigTurn, this.curSmallTurn, this.curTurnEntity())
	msg := usercmd.TurnInfoS2CMsg{
		BigTurnIndex:      this.curBigTurn,
		SmallTurnIndex:    this.curSmallTurn,
		CurEntityPosIndex: this.curTurnEntity(),
	}
	data, err := util.EncodeCmd(usercmd.UserCmd_TurnInfo, &msg)
	if err != nil {
		return
	}
	this.turnRoom.BroadcastMsg(data)
}
