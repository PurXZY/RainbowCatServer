package idmgr

import (
	"sync"
	"sync/atomic"
)

type IdMgr struct {
	uniqId     uint32
}

var (
	mgr     *IdMgr
	mgrOnce sync.Once
)

func GetMe() *IdMgr {
	if mgr == nil {
		mgrOnce.Do(func() {
			mgr = &IdMgr{
				uniqId:     0,
			}
		})
	}
	return mgr
}

func (this *IdMgr) GenUniqId() uint32 {
	atomic.AddUint32(&this.uniqId, 1)
	return this.uniqId
}


