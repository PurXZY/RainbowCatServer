package main

import (
	"base/log"
	"gamserver"
	"os"
)

func main() {
	log.InitLog(os.Stdout, os.Stdout, os.Stdout)
	ser := gamserver.NewGameServer()
	ser.Start()
	log.Info.Println("server all over")
}
