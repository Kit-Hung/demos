/*
 * @Author:       Kit-Hung
 * @Date:         2024/6/19 18:01
 * @Description： grpc client
 */
package main

import (
	"context"
	"fmt"
	pb "github.com/Kit-Hung/demos/golang/13-grpc/client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// 连接到 server（使用证书）
	// creds, _ := credentials.NewClientTLSFromFile("xxx.pem", "*.xx.com")
	// conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(creds))

	// 连接到 server
	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect the grpc server: %v", err)
		return
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			log.Fatalf("failed to close the grpc connection: %v", err)
		}
	}(conn)

	// 建立连接
	client := pb.NewSayHelloClient(conn)
	// 执行 rpc 调用
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "kit"})
	if err != nil {
		log.Fatalf("failed to call SayHello func: %v", err)
		return
	}
	fmt.Println(resp.ResponseMsg)
}
