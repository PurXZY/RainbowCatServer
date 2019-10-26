package main

import (
	"base/util"
	"client/userclient"
	"time"
)

func main() {
	cli := userclient.NewTestClient()
	cli.ConnectServer("127.0.0.1:8888")

	ticker := time.NewTicker(1 * time.Microsecond)

	for _ = range ticker.C {
		cli.GetTcpTask().SendData(util.IntToBytes(time.Now().Second()))
	}
}
