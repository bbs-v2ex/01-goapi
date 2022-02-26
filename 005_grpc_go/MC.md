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