package battleentity

import (
	"usercmd"
)

type BattleEntity struct {
	posIndex   usercmd.PosIndex
	entityType uint32
	Prop       EntityProp
}

func NewBattleEntity(posIndex usercmd.PosIndex, entityType uint32) *BattleEntity {
	entity := &BattleEntity{
		posIndex:   posIndex,
		entityType: entityType,
	}
	return entity
}

func (this *BattleEntity) GetType() uint32 {
	return this.entityType
}
