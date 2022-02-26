package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"

	proto2 "n01/005_grpc_go/proto"
)

func main() {
	req := proto2.HelloRequest{
		Name: "bobby",
	}
	//编码
	rsp, _ := proto.Marshal(&req)
	fmt.Println(string(rsp))
	var r2 proto2.HelloRequest
	//解码
	err := proto.Unmarshal(rsp, &r2)
	if err != nil {
		return
	}
	fmt.Println(r2)
}
