package net

import (
	"base/log"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"syscall"
)

type IService interface {
	Init() bool
	MainLoop()
	Final()
}

type Service struct {
	userService IService
	Server      *TcpServer
	terminate   bool
}

func (this *Service) SetUserService (service IService) {
	this.userService = service
}

func (this *Service) Terminate() {
	this.terminate = true
	_ = this.Server.Close()
}

func (this *Service) isTerminate() bool {
	return this.terminate
}

func (this *Service) SetCpuNum(num int) {
	if num > 0 {
		runtime.GOMAXPROCS(num)
	} else if num == -1 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
}

func (this *Service) Start() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Error.Println("err: ", err, "\n", string(debug.Stack()))
		}
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGPIPE, syscall.SIGHUP)
	go func() {
		for sig := range ch {
			log.Info.Println("receive signal: ", sig)
			switch sig {
			case syscall.SIGPIPE:
				log.Error.Println("SIGPIPE")
			default:
				this.Terminate()
			}
		}
	}()

	runtime.GOMAXPROCS(runtime.NumCPU())

	if !this.userService.Init() {
		return false
	}

	for !this.isTerminate() {
		this.userService.MainLoop()
	}

	this.userService.Final()
	return true
}

