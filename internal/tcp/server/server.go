package server

import (
	"bufio"
	"go.uber.org/zap"
	"net"
)


type Server interface{
	Listen(address string) error
	Start()
	Receive()
	BroadCast()
}
var _ Server = (*TcpServer)(nil)

type TcpServer struct {
	logger *zap.Logger
	listener net.Listener
	clients []*Client
}


func NewTcpServer()  *TcpServer{
	return &TcpServer{}
}

func (t *TcpServer) Listen(address string) error{
	l, err := net.Listen("tcp",address)
	if err==nil {
		t.listener = l
	}
	return err
}

func (t *TcpServer) Start() {
	for{
		conn,err := t.listener.Accept()
		if err!=nil {

		}
		str,err:= bufio.NewReader(conn).ReadString('\n')
	}


}

func (t *TcpServer) Receive() {
	panic("implement me")
}

func (t *TcpServer) BroadCast() {
	panic("implement me")
}

type Client struct {
	name string
	conn net.Conn
}




