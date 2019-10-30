package sessionmgr

import (
	"gamserver/entity/session"
	"gamserver/mgr/idmgr"
	"net"
	"sync"
)

type SessionMgr struct {
	sessions map[uint32]*session.Session
}

var (
	mgr     *SessionMgr
	mgrOnce sync.Once
)

func GetMe() *SessionMgr {
	if mgr == nil {
		mgrOnce.Do(func() {
			mgr = &SessionMgr{
				sessions: make(map[uint32]*session.Session),
			}
		})
	}
	return mgr
}

func (this *SessionMgr) AddNewSession(conn net.Conn) {
	uniqId := idmgr.GetMe().GenUniqId()
	target := session.NewSession(conn, uniqId)
	this.sessions[uniqId] = target
	target.Start()
}