package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	var reply string
	err = client.Call("HelloService.Hello", "body", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
