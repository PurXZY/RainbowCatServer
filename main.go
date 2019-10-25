package main

import (
	"base/log"
	"base/net"
	"os"
)

func main() {
	log.InitLog(os.Stdout, os.Stdout, os.Stdout)
	log.Info.Println("test")
	log.Debug.Printf("test %d", 1)
	log.Error.Println("asdd")

	server := net.Server{}
}
