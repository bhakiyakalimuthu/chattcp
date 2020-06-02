package main

import (
	"go.uber.org/zap"

	"github.com/bhakiyakalimuthu/chattcp/internal/tcp/client"
)

func main()  {

	c:= client.NewTcpClient(zap.NewNop(),nil,"client1")
	c.Dial(":6666")
	c.Send()
}