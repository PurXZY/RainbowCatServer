package netmsgdispatcher

import (
	"base/log"
	"usercmd"
)

// 网络消息处理器
type MsgHandler func(data []byte)

type NetMsgDispatcher struct {
	handlerMap map[usercmd.UserCmd]MsgHandler
}

func (this *NetMsgDispatcher) Init() {
	this.handlerMap = make(map[usercmd.UserCmd]MsgHandler)
}

func (this *NetMsgDispatcher) RegisterHandler(cmd usercmd.UserCmd, cb MsgHandler) {
	this.handlerMap[cmd] = cb
}

func (this *NetMsgDispatcher) Call(cmd usercmd.UserCmd, data []byte) {
	cb, ok := this.handlerMap[cmd]
	if ok {
		cb(data)
	} else {
		log.Error.Println("NetMsgDispatcher Call unknown cmd:", cmd)
	}
}

