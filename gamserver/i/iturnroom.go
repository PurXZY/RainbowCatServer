package i

import "usercmd"

type ITurnRoom interface {
	BroadcastMsg(data []byte)
	GetAllEntitiesSpeed() map[usercmd.PosIndex]uint32
}
