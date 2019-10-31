package turnroom

import (
	"base/util"
	"gamserver/entity/battleentity"
	"gamserver/i"
	"usercmd"
)

type EntityMgr struct {
	turnRoom       i.ITurnRoom
	battleEntities map[uint32]*battleentity.BattleEntity
}

func NewEntityMgr() *EntityMgr {
	eMgr := &EntityMgr{
		battleEntities: make(map[uint32]*battleentity.BattleEntity),
	}
	return eMgr
}

func (this *EntityMgr) Init(room i.ITurnRoom) {
	this.turnRoom = room
	this.initBattleEntities()
	this.notifyClientAllBattleEntities()
}

func (this *EntityMgr) initBattleEntities() {
	enemyData := map[usercmd.PosIndex]uint32{
		usercmd.PosIndex_PosELeft:   1,
		usercmd.PosIndex_PosECenter: 1,
		usercmd.PosIndex_PosERight:  2,
	}
	for posIndex, entityType := range enemyData {
		entity := battleentity.NewBattleEntity(uint32(posIndex), entityType)
		this.battleEntities[uint32(posIndex)] = entity
	}

	myData := map[usercmd.PosIndex]uint32{
		usercmd.PosIndex_PosBCenter: 3,
	}
	for posIndex, entityType := range myData {
		entity := battleentity.NewBattleEntity(uint32(posIndex), entityType)
		this.battleEntities[uint32(posIndex)] = entity
	}
}

func (this *EntityMgr) notifyClientAllBattleEntities() {
	msg := usercmd.CreateAllBattleEntitiesS2CMsg{}
	for _, entity := range this.battleEntities {
		msg.Entities = append(msg.Entities, entity.Prop.GetPropData())
	}
	data, err := util.EncodeCmd(usercmd.UserCmd_CreateAllBattleEntities, &msg)
	if err != nil {
		return
	}
	this.turnRoom.BroadcastMsg(data)
}

func (this *EntityMgr) GetAllBattleEntities() []*battleentity.BattleEntity {
	v := make([]*battleentity.BattleEntity, len(this.battleEntities))
	for _, value := range this.battleEntities {
		v = append(v, value)
	}
	return v
}

func (this *EntityMgr) GetAllEntitiesSpeed() map[uint32]uint32 {
	v := make(map[uint32]uint32, len(this.battleEntities))
	for key, value := range this.battleEntities {
		v[key] = value.Prop.MoveSpeed
	}
	return v
}