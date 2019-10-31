package battleentity

import (
	"gamserver/i"
	"gamserver/mgr/datamgr"
	"usercmd"
)

type EntityProp struct {
	owner          i.IBattleEntity
	Health         uint32
	PhysicalAttack uint32
	MagicAttack    uint32
	PhysicalDefend uint32
	MagicDefend    uint32
	MoveSpeed      uint32
}

func (this *EntityProp) Init(entity i.IBattleEntity, u uint32) {
	this.owner = entity
	data := datamgr.GetMe().GetPropData(u)
	if data == nil {
		return
	}
	this.Health = data["Health"]
	this.PhysicalAttack = data["PhysicalAttack"]
	this.MagicAttack = data["MagicAttack"]
	this.PhysicalAttack = data["PhysicalAttack"]
	this.MagicDefend = data["MagicDefend"]
	this.MoveSpeed = data["MoveSpeed"]
}

func (this *EntityProp) GetPropData() *usercmd.BattleEntity {
	e := &usercmd.BattleEntity{
		PosIndex:       this.owner.GetPosIndex(),
		EntityType:     this.owner.GetType(),
		Health:         this.Health,
		PhysicalAttack: this.PhysicalAttack,
		MagicAttack:    this.MagicAttack,
		PhysicalDefend: this.PhysicalDefend,
		MagicDefend:    this.MagicDefend,
		MoveSpeed:      this.MoveSpeed,
	}
	return e
}
