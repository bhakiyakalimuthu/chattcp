package main

import (
	"go.uber.org/zap"

	"github.com/bhakiyakalimuthu/chattcp/internal/tcp/server"
	"github.com/bhakiyakalimuthu/chattcp/pkg/tcp/client"
)

func main()  {
	s:= server.NewTcpServer(zap.NewNop(),nil,nil)
	s.Listen(":6666")
	s.Start()

	c:= client.NewTcpClient(zap.NewNop(),nil,"client1")
	c.Dial(":6666")
	c.Send()
}