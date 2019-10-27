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
				randomValue := rand.Int31()
				sendData := util.Uint32ToBytes(uint32(randomValue))
				cli.GetTcpTask().SendData(sendData)
				log.Debug.Println("send data:", sendData, "value:", randomValue)
			case <- out:
				break Loop
		}
	}
	time.Sleep(3 * time.Second)
	log.Info.Println("client all over")
}
