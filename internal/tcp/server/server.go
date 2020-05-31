package server

import (
	"bufio"
	"go.uber.org/zap"
	"net"
	"time"
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

func NewTcpServer(logger *zap.Logger, listener net.Listener, clients []*Client)  *TcpServer{
	return &TcpServer{
		logger:   logger,
		listener: listener,
		clients:  clients,
	}
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
			t.logger.Error("failed to accept connection")
		}
		str,err:= bufio.NewReader(conn).ReadString('\n')
		if err!=nil {
			t.logger.Error("failed to read from connection")
		}
		t.logger.Info(str)
		t:= time.Now().Format(time.RFC3339)
		conn.Write([]byte(t))
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




