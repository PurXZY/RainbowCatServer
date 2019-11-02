package avatar

import (
	"gamserver/mgr/turnroommgr"
	"gamserver/netmsgdispatcher"
	"usercmd"
)

type AvatarNetMsgDispatcher struct {
	dispatcher         netmsgdispatcher.NetMsgDispatcher
	selfAvatar         *Avatar
}

func (this *AvatarNetMsgDispatcher) Init(avatar *Avatar){
	this.selfAvatar = avatar
	this.dispatcher.Init()
	this.RegMsgHandler()
}

func(this *AvatarNetMsgDispatcher) RegMsgHandler() {
	this.dispatcher.RegisterHandler(usercmd.UserCmd_IntoRoomReq, this.OnReqIntoRoom)
	this.dispatcher.RegisterHandler(usercmd.UserCmd_CastOperationReq, this.OnReqCastOperation)
}

func (this *AvatarNetMsgDispatcher) OnReqIntoRoom(data []byte) {
	turnroommgr.GetMe().AddNewTurnRoom(this.selfAvatar)
}

func (this *AvatarNetMsgDispatcher) OnReqCastOperation(data []byte) {
	//recvMsg, ok := util.DecodeCmd(data, &usercmd.CastOperationC2SMsg{}).(*usercmd.CastOperationC2SMsg)
	//if !ok {
	//	return
	//}
}