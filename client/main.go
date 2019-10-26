package main

import (
	"base/log"
	"base/util"
	"client/userclient"
	"math/rand"
	"os"
	"time"
)

func main() {
	log.InitLog(os.Stdout, os.Stdout, os.Stdout)
	cli := userclient.NewTestClient()
	cli.ConnectServer("127.0.0.1:8888")

	c := time.Tick(1 * time.Second)
	out := time.After(5 * time.Second)

	Loop:
	for {
		select {
			case <- c:
				randomValue := rand.Int()
				cli.GetTcpTask().SendData(util.Uint32ToBytes(uint32(randomValue)))
				log.Debug.Println("send value:", randomValue)
			case <- out:
				break Loop
		}
	}

	log.Info.Println("client all over")
}
