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
