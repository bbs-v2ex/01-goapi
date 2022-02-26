package main

import (
	"context"
	"fmt"
	"n01/007_grpc_stream/proto"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := proto.NewGreeterClient(conn)
	////001 接收服务端发送的流
	//res, err := client.GetStream(context.Background(), &proto.StreamReqData{
	//	Data: "love imooc",
	//})
	//if err != nil {
	//	panic(err)
	//}
	//for {
	//	recv, err := res.Recv()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(recv)
	//}
	//// 客户端流模式
	//putS, _ := client.PutStream(context.Background())
	//i := 0
	//for {
	//	putS.Send(&proto.StreamReqData{
	//		Data: fmt.Sprintf("%d", i),
	//	})
	//	i++
	//	if i > 10 {
	//		break
	//	}
	//}
	// 双向
	allStr, _ := client.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息：" + data.Data)
		}
	}()

	//1. 集中学习protobuf， grpc

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamReqData{Data: "慕课网"})
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}
