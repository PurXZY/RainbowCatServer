package avatar

import (
	"gamserver/mgr/turnroommgr"
	"gamserver/netmsgdispatcher"
	"usercmd"
)

type AvatarNetMsgDispatcher struct {
	dispatcher netmsgdispatcher.NetMsgDispatcher
	selfAvatar *Avatar
}

func (this *AvatarNetMsgDispatcher) Init(avatar *Avatar){
	this.selfAvatar = avatar
	this.dispatcher.Init()
	this.RegMsgHandler()
}

func(this *AvatarNetMsgDispatcher) RegMsgHandler() {
	this.dispatcher.RegisterHandler(usercmd.UserCmd_IntoRoomReq, this.OnReqIntoRoom)
}

func (this *AvatarNetMsgDispatcher) OnReqIntoRoom(data []byte) {
	turnroommgr.GetMe().AddNewTurnRoom(this.selfAvatar)
}