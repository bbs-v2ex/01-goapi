package main

import (
	"net"
	"net/rpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = request + *reply
	return nil
}

func main() {
	//创建链接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
		return
	}
	//注册
	_ = rpc.RegisterName("HelloService", &HelloService{})
	for {
		conn, _ := listener.Accept() //接收请求
		rpc.ServeConn(conn)
	}
}
