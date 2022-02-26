package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}
func main() {
	//实例化一个server
	listener, _ := net.Listen("tcp", ":1234")

	//注册处理逻辑
	_ = rpc.RegisterName("HelloService", &HelloService{})

	//启动服务
	for {
		conn, _ := listener.Accept() //当链接进来时
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

	//rpc.ServeConn(conn)
}
