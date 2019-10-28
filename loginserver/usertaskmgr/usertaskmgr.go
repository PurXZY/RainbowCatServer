package usertaskmgr

import (
	"base/log"
	"base/net"
	"math/rand"
	"sync"
)

type UserTaskMgr struct {
	mutex sync.RWMutex
	tasks map[uint32]net.ITcpTask
}

var (
	userTaskMgr     *UserTaskMgr
	userTaskMgrOnce sync.Once
)

func GetMe() *UserTaskMgr {
	if userTaskMgr == nil {
		userTaskMgrOnce.Do(func() {
			userTaskMgr = &UserTaskMgr{
				mutex: sync.RWMutex{},
				tasks: make(map[uint32]net.ITcpTask),
			}
		})
	}
	return userTaskMgr
}

func(this *UserTaskMgr) AddNewPlayerTask(task net.ITcpTask) bool {
	taskId := rand.Uint32() % 10000
	if this.GetPlayerTask(taskId) != nil {
		log.Error.Println("same taskId:", taskId)
		return false
	}
	this.mutex.Lock()
	this.tasks[taskId] = task
	this.mutex.Unlock()
	task.SetId(taskId)
	return true
}

func(this *UserTaskMgr) DeletePlayerTask(id uint32) {
	task := this.GetPlayerTask(id)
	if task == nil {
		log.Error.Println("DeletePlayerTask no task")
		return
	}
	this.mutex.Lock()
	delete(this.tasks, id)
	this.mutex.Unlock()
	log.Debug.Println("DeletePlayerTask remove taskId:", id)
}

func(this *UserTaskMgr) GetPlayerTask(id uint32) net.ITcpTask {
	defer this.mutex.RUnlock()
	this.mutex.RLock()
	task, ok := this.tasks[id]
	if !ok {
		return nil
	}
	return task
}