# 实现 grpc 通信的 demo
## 总体思路
1. 安装依赖
2. 编写 proto 文件
3. 生成相关代码
4. 实现具体逻辑
5. 服务端： 启动 grpc server
6. 客户端： 连接 grpc server 并执行 rpc 调用




## 安装依赖
1. 安装 protoc
   [protobuf](https://github.com/protocolbuffers/protobuf)
   
2. 安装相关包
   ```shell
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```
   
3. 配置环境变量
   ```shell
   export PATH="$PATH:$(go env GOPATH)/bin"
   ```
4. 


## 编写 proto 文件

1. 设置版本
   ```protobuf
   // 指定使用 proto3
   syntax = "proto3";
   ```
      
2. 定义代码生成的目录和包
   ```protobuf
   // 最后生成的 go 文件属于哪个目录哪个包
   // . 代表在当前目录生成
   // service 代表生成的 go 文件的包名是 service
   option go_package = ".;service";
   ```
      
3. 定义服务
   ```protobuf
   // 定义服务 SayHello ，在这个服务中有一个 rpc 方法，接收客户端参数 HelloRequest ，在返回服务端响应 HelloResponse
   service SayHello {
       rpc SayHello(HelloRequest) returns (HelloResponse) {}
   }
   ```
      
4. 定义参数结构
   ```protobuf
   // message 相当于 go 中的结构体
   // repeated 会生成切片
   message HelloRequest {
       string requestName = 1;
       repeated int64 nums = 2;
   }
      
   message HelloResponse {
       string responseMsg = 1;
   }
   ```


## 生成对应的代码
   ```shell
   protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative .\xx.proto
   ```


## 实现具体逻辑
```go
// Server 创建具体的实现结构体
type Server struct {
	pb.UnimplementedSayHelloServer
}

// SayHello 根据实际需求实现对应的方法
func (s *Server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}
```

## 服务端： 启动 grpc server
```go
// 开启端口
listen, err := net.Listen("tcp", ":8888")
if err != nil {
    log.Fatalf("Failed to open port: %v", err)
    return
}

// 创建 grpc 服务
grpcServer := grpc.NewServer()
// 在 grpc 服务端中注册服务
pb.RegisterSayHelloServer(grpcServer, &Server{})

// 启动服务
err = grpcServer.Serve(listen)
if err != nil {
    log.Fatalf("Failed to serve grpc server: %v", err)
    return
}
```


## 客户端： 连接 grpc server 并执行 rpc 调用
```go

```