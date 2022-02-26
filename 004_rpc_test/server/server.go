package main

import (
	"n01/004_rpc_test/hanlder"
	"net"
	"net/rpc"

	"n01/004_rpc_test/server_proy"
)

func main() {
	//创建链接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
		return
	}
	//注册
	//_ = rpc.RegisterName(hanlder.HelloServiceName, &hanlder.HelloService{})
	server_proy.RegisterHelloService(hanlder.HelloService{})
	for {
		conn, _ := listener.Accept() //接收请求
		go rpc.ServeConn(conn)
	}
}
