package avatarmgr

import (
	"gamserver/entity/avatar"
	"gamserver/i"
	"gamserver/mgr/idmgr"
	"sync"
)

type AvatarMgr struct {
	avatars map[uint32]*avatar.Avatar
}

var (
	mgr     *AvatarMgr
	mgrOnce sync.Once
)

func GetMe() *AvatarMgr {
	if mgr == nil {
		mgrOnce.Do(func() {
			mgr = &AvatarMgr{
				avatars: make(map[uint32]*avatar.Avatar),
			}
		})
	}
	return mgr
}

func (this *AvatarMgr) AddNewAvatar(session i.ISession, loginName string) {
	uniqId := idmgr.GetMe().GenUniqId()
	target := avatar.NewAvatar(uniqId, session, loginName)
	this.avatars[uniqId] = target
}