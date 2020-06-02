package main

import (
	"go.uber.org/zap"

	"github.com/bhakiyakalimuthu/chattcp/internal/tcp/server"
)

func main()  {
	s:= server.NewTcpServer(zap.NewNop(),nil,nil)
	s.Listen(":6666")
	s.Start()

}