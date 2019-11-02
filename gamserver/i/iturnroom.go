package i

type ITurnRoom interface {
	BroadcastMsg(data []byte)
	GetAllEntitiesSpeed() map[uint32]uint32
}
