package turnroom

import (
	"gamserver/i"
	"sort"
	"usercmd"
)

type TurnLogic struct {
	turnRoom     i.ITurnRoom
	curBigTurn   uint32
	curSmallTurn uint32
	entityTurnInfo []usercmd.PosIndex
}

func (this *TurnLogic) Init(room i.ITurnRoom) {
	this.turnRoom = room
	this.curBigTurn = 1
	this.curSmallTurn = 0
}

type PosSpeedPair struct {
	pos   usercmd.PosIndex
	speed uint32
}
type PairList []PosSpeedPair
func(pair PairList) Swap(i, j int) {}
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].speed < p[j].speed }

func (this *TurnLogic) SortEntitySpeedInfo() {
	speedMap := this.turnRoom.GetAllEntitiesSpeed()
	plist := make(PairList, len(speedMap))
	for posIndex, speed := range speedMap {
		plist = append(plist, PosSpeedPair{
			pos:   posIndex,
			speed: speed,
		})
	}
	sort.Sort(plist)
	this.entityTurnInfo = this.entityTurnInfo[:0]
	for _, pair := range plist{
		this.entityTurnInfo = append(this.entityTurnInfo, pair.pos)
	}
}
