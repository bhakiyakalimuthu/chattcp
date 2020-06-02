package client

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"net"
	"os"
)

type Client interface{

	Dial(address string) error
	Send()
	Close()
	Receive()
}
var _ Client = (*TcpClient)(nil)

type TcpClient struct {
	logger *zap.Logger
	conn net.Conn
	name string
}

func NewTcpClient(logger *zap.Logger, conn net.Conn, name string)  *TcpClient {
	return &TcpClient{
		logger: logger,
		conn:   conn,
		name:   name,
	}
}
func (t *TcpClient) Dial(address string) error{
	conn,err := net.Dial("tcp",address)
	if err==nil {
		t.conn = conn
	}
	t.logger.Error("failed to dial",zap.String("client name",t.name))
	return err
}

func (t *TcpClient) Send() {
	reader:= bufio.NewReader(os.Stdin)
	input,err := reader.ReadString('\n')
	if err!=nil {
		t.logger.Error("failed to read from user input", zap.Error(err))
	}
	t.logger.Info("getting input from user",zap.String("input",input))
	fmt.Fprintf(t.conn,input)
}

func (t *TcpClient) Close() {
	panic("implement me")
}

func (t *TcpClient) Receive() {
	panic("implement me")
}
