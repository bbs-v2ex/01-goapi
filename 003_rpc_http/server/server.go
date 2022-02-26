package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}
func main() {

	_ = rpc.RegisterName("HelloService", &HelloService{})

	//实例化一个server , 使用http 的URL 处理
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	//启动http 服务器
	http.ListenAndServe(":1234", nil)

	//http.ListenAndServe()
	//listener, _ := http.("tcp", ":1234")
	//
	////注册处理逻辑
	//_ = rpc.RegisterName("HelloService", &HelloService{})
	//
	////启动服务
	//for {
	//	conn, _ := listener.Accept() //当链接进来时
	//	rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	//}
}
