package i

type ISession interface {
	SetOwner(ISessionOwner)
	SendData([]byte)
	ParseMsg([]byte) bool
	OnClose()
}
