package battleentity

type BattleEntity struct {
	posIndex   uint32
	entityType uint32
	Prop       EntityProp
}

func NewBattleEntity(posIndex uint32, entityType uint32) *BattleEntity {
	entity := &BattleEntity{
		posIndex:   posIndex,
		entityType: entityType,
	}
	entity.Prop.Init(entity, entityType)
	return entity
}

func (this *BattleEntity) GetType() uint32 {
	return this.entityType
}

func (this *BattleEntity) GetPosIndex() uint32 {
	return this.posIndex
}