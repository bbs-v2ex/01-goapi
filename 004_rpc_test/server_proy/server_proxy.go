package server_proy

import (
	"n01/004_rpc_test/hanlder"
	"net/rpc"
)

type HelloServicer interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv hanlder.HelloService) error {
	return rpc.RegisterName(hanlder.HelloServiceName, srv)

}
