package net

import (
	"base/log"
	"net"
)

type Server struct {
	listener *net.TCPListener
}

func (this *Server) Bind(address string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		log.Error.Println("bind fail ", address)
		return
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Error.Println("listen fail ", address)
		return
	}

	log.Info.Println("bind success ", address)
	this.listener = listener
}
