package main

import (
	"base/log"
	"base/net"
	"base/util"
	"client/userclient"
	"os"
	"time"
	"usercmd"
)

func main() {
	log.InitLog(os.Stdout, os.Stdout, os.Stdout)
	cli := userclient.NewTestClient()
	ret := cli.ConnectServer("127.0.0.1:8888")
	if !ret {
		return
	}

	c := time.Tick(1 * time.Second)
	out := time.After(5 * time.Second)

	Loop:
	for {
		select {
			case <- c:
				SendProtoData(cli.GetTcpTask())
			case <- out:
				break Loop
		}
	}
	time.Sleep(1 * time.Second)
	cli.CloseConnection()
	log.Info.Println("client all over")
}

func SendProtoData(task net.ITcpTask) {
	msg := usercmd.LoginC2SMsg{
		Name:"xzy",
	}
	data, _ := util.EncodeCmd(usercmd.UserCmd_LoginReq, &msg)
	task.SendData(data)
}