package userclient

import (
	"base/net"
	"client/usertask"
)
import engineNet "net"

type TestClient struct {
	net.TcpClient
}

func NewTestClient() *TestClient {
	cli := &TestClient{}
	cli.SetUserClient(cli)
	return cli
}

func (this *TestClient) OnConnected(conn engineNet.Conn) net.ITcpTask {
	task := usertask.NewTask(conn)
	task.Start()
	this.SetTcpTask(task)
	return task
}
