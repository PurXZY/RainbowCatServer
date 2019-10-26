package net

import (
	"base/log"
	"net"
)

type ITcpClient interface {
	OnConnected(conn net.Conn) ITcpTask
}

type TcpClient struct {
	userClient ITcpClient
	userTask ITcpTask
}

func (this *TcpClient) SetUserClient (userClient ITcpClient) {
	this.userClient = userClient
}

func (this *TcpClient) ConnectServer(address string) bool {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		log.Error.Println("resolve fail address: ", address)
		return false
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Error.Println("dial fail address: ", address)
		return false
	}
	_ = conn.SetKeepAlive(true)
	_ = conn.SetNoDelay(true)
	_ = conn.SetWriteBuffer(64 * 1024)
	_ = conn.SetReadBuffer(64 * 1024)
	this.userClient.OnConnected(conn)
	return true
}

func (this *TcpClient) SetTcpTask(task ITcpTask) {
	this.userTask = task
}

func (this *TcpClient) GetTcpTask() ITcpTask {
	return this.userTask
}
