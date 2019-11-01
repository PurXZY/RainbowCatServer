package i

import "usercmd"

type IOwner interface {
	SendData(data []byte)
}

type ISessionOwner interface {
	IOwner
	GetName() string
	OnRecvMsg(usercmd.UserCmd, []byte)
}

type IRoomOwner interface {
	IOwner
	GetName() string
	SetTurnRoom(ITurnRoom)
}

type IBattleEntity interface {
	GetPosIndex() uint32
	GetType() uint32
}