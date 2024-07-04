/*
 * @Author:       Kit-Hung
 * @Date:         2024/6/19 18:01
 * @Description： grpc server
 */
package main

import (
	"context"
	pb "github.com/Kit-Hung/demos/golang/13-grpc/server/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

// Server 创建具体的实现结构体
type Server struct {
	pb.UnimplementedSayHelloServer
}

// SayHello 根据实际需求实现对应的方法
func (s *Server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{ResponseMsg: "hello " + req.RequestName}, nil
}

func main() {
	// 开启端口
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Failed to open port: %v", err)
		return
	}

	// 创建 grpc 服务
	grpcServer := grpc.NewServer()

	// 创建 grpc 服务（使用证书）
	// creds, _ := credentials.NewServerTLSFromFile("xx.pem", "xx.key")
	// grpcServer := grpc.NewServer(grpc.Creds(creds))

	// 在 grpc 服务端中注册服务
	pb.RegisterSayHelloServer(grpcServer, &Server{})

	// 启动服务
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to serve grpc server: %v", err)
		return
	}
}
