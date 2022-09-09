package main

import (
	"context"
	"fmt"
	client2 "github.com/jayson-hu/rpc/grpc/middleware/client"
	"github.com/jayson-hu/rpc/grpc/middleware/server"
	"google.golang.org/grpc/metadata"
	"time"

	"github.com/jayson-hu/rpc/grpc/simple/server/pb"
	"google.golang.org/grpc"
)

func main() {

	//第二中：添加认证方式
	credital :=client2.NewAuthenication("admin", "123456")


	conn, err := grpc.DialContext(context.Background(),"localhost:1234",
		grpc.WithInsecure(),grpc.WithPerRPCCredentials(credital))
	if err != nil {
		panic(err)
	}
	//grpc生成一个client
	client := pb.NewHelloServiceClient(conn)
	//req <--> resp
	//携带凭证改进版
	//添加认证凭证信息
	crendential := server.NewClientCredential("admin","123456")// 第一种方式
	ctx := metadata.NewOutgoingContext(context.Background(),crendential) // 第一种方式

	// 第二种方式： 使用client auth
	//ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs())


	//crendential2 := server.NewClientCredential("admin","1234567")
	//ctx2 := metadata.NewOutgoingContext(context.Background(),crendential2)
	//
	response, err := client.Hello(ctx, &pb.Request{Value: "alice"})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Value)

	//stream
	//stream, err := client.Channel(context.Background())
	//流式处理
	stream, err := client.Channel(ctx)
	if err != nil {
		panic(err)
	}
	// 启用 goroutine 发送请求
	go func() {
		for {
			err := stream.Send(&pb.Request{
				Value: "alice",
			})
			if err != nil {
				panic(err)
			}
			time.Sleep(time.Second * 1)
		}
	}()

	//主进程负责接收请求相应
	for {
		resp, err := stream.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.Value)

	}

}
