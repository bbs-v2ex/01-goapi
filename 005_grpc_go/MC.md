首先下载 protocol 

```bash
https://github.com/protocolbuffers/protobuf
```

![image-20220226155731029](https://gitee.com/ASeditor_admin/typora_img/raw/master/img/image-20220226155731029.png)

下载这个 解压 然后添加环境变量

![image-20220226160055913](https://gitee.com/ASeditor_admin/typora_img/raw/master/img/image-20220226160055913.png)

![image-20220226160138570](https://gitee.com/ASeditor_admin/typora_img/raw/master/img/image-20220226160138570.png)

![image-20220226160241936](https://gitee.com/ASeditor_admin/typora_img/raw/master/img/image-20220226160241936.png)

检查环境变量是否配置成功

![image-20220226160319587](https://gitee.com/ASeditor_admin/typora_img/raw/master/img/image-20220226160319587.png)

下载 go 的依赖包

```
go get github.com/golang/protobuf/protoc-gen-go
```

```protobuf
syntax = "proto3";
option go_package = ".;proto";
//grpc 接口
service Greeter {
  //方法
  rpc SayHello (HelloRequest) returns (HelloReply);
}
message HelloRequest {
  string  name = 1;
}
message HelloReply {
  string message = 1;
}
```

protoc 可以通过文件生成各个语言版本

## Protocol

生成 go 语言版本

```protobuf
protoc -I . helloworld.proto --go_out=plugins=grpc:.
```

```go
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

```

![image-20220226164808672](https://gitee.com/ASeditor_admin/typora_img/raw/master/img/image-20220226164808672.png)

对比之下, prptocol 相对于 json 拥有更高的压缩比,数据更小

### 一元调用

Protocol

```protobuf
syntax = "proto3";
//生成文件的 go 的包名称
option go_package = ".;proto";
//grpc 接口
service Greeter {
  //方法
  rpc SayHello (HelloRequest) returns (HelloReply);
}
message HelloRequest {
  string  name = 1;
}
message HelloReply {
  string message = 1;
}
```

server

```go
package main

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"n01/006_grpc_go_2/proto"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	//实例化 grpc
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	//启动服务
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	err = g.Serve(listener)
	if err != nil {
		panic(err)
	}
}

```

client

```go
package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"n01/006_grpc_go_2/proto"
)

func main() {
	//conn, _ := grpc.Dial("localhost:8080", grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "bobby"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r)

}

```

![image-20220226175126481](https://gitee.com/ASeditor_admin/typora_img/raw/master/img/image-20220226175126481.png)

### 流模式

可以源源不断的传输数据,适合传输一系大数据,适合B/C长时间连接

![image-20220226175422780](https://gitee.com/ASeditor_admin/typora_img/raw/master/img/image-20220226175422780.png)

![image-20220226175457160](https://gitee.com/ASeditor_admin/typora_img/raw/master/img/image-20220226175457160.png)

![image-20220226175508145](https://gitee.com/ASeditor_admin/typora_img/raw/master/img/image-20220226175508145.png)

![image-20220226175656952](https://gitee.com/ASeditor_admin/typora_img/raw/master/img/image-20220226175656952.png)

![image-20220226175741678](https://gitee.com/ASeditor_admin/typora_img/raw/master/img/image-20220226175741678.png)

client

```go
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

```

server

```go
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

```

proto

```protobuf
syntax = "proto3";
option go_package = ".;proto";


service Greeter {
  rpc GetStream (StreamReqData) returns (stream StreamResData); //服务端流模式
  rpc PutStream (stream StreamReqData) returns (StreamResData); //客户端流模式
  rpc AllStream (stream StreamReqData) returns (stream StreamResData); //双向流模式
}

message StreamReqData {
  string data = 1;
}
message StreamResData {
  string data = 1;
}
```

