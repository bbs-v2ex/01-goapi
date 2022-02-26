package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"

	"n01/007_grpc_stream/proto"
)

const PORT = ":50052"

type Server struct {
}

// 001 向客户端发送流
func (s *Server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%d-%d", i, time.Now().UnixNano()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
		i++
	}
	return nil
}

func (s *Server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	for {
		recv, err := cliStr.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(recv)
	}
	return nil
}

func (s *Server) AllStream(allStr proto.Greeter_AllStreamServer) error {

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息：" + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamResData{Data: "我是服务器"})
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	return nil

	//wg := sync.WaitGroup{}
	//wg.Add(2)
	//go func() {
	//	defer wg.Done()
	//	recv, err := aliStr.Recv()
	//	if err != nil {
	//		return
	//	}
	//	fmt.Println("收到客户端消息", recv)
	//}()
	//
	//go func() {
	//	i := 0
	//	for {
	//		i++
	//		if i > 10 {
	//			break
	//		}
	//		defer wg.Done()
	//		_ = aliStr.Send(&proto.StreamResData{
	//			Data: "向客户端发送消息",
	//		})
	//	}
	//
	//	time.Sleep(time.Second)
	//}()
	//return nil
}

//func (s *Server) GetStream (ctx context.Context, req proto.StreamReqData) (*proto.StreamResData, error) {
//	return nil,nil
//}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &Server{})
	s.Serve(lis)
}
