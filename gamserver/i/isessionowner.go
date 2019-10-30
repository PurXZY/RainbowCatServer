package i

type IOwner interface {
	SendData(data []byte)
}

type ISessionOwner interface {
	IOwner
	GetName() string
	ReqIntoRoom()
}
