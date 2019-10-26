package main

import (
	"base/log"
	"loginserver/userservice"
	"os"
)

func main() {
	log.InitLog(os.Stdout, os.Stdout, os.Stdout)
	ser := userservice.NewLoginServer()
	ser.Start()
	log.Info.Println("all over")
}
