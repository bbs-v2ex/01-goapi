package main

import (
	"fmt"

	"n01/004_rpc_test/client_proxy"
)

func main() {
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")
	//只想写业务代码,不想关注每个函数的名称
	var reply string
	err := client.Hello("body", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
	//client, err := rpc.Dial("tcp", "localhost:1234")
	//if err != nil {
	//	panic(err)
	//}
	//var reply string
	//err = client.Call(hanlder.HelloServiceName+".Hello", "body", &reply)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(reply)
}
