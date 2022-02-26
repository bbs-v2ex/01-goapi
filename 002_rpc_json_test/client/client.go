package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	err = client.Call("HelloService.Hello", "body", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
