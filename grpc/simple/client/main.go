package main

import (
	"context"
	"fmt"
	"github.com/jayson-hu/rpc/grpc/simple/server/pb"
	"google.golang.org/grpc"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	//grpc生成一个client
	client := pb.NewHelloServiceClient(conn)
	//req <--> resp
	response, err := client.Hello(context.Background(), &pb.Request{
		Value: "alice",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Value)

	//stream
	stream, err := client.Channel(context.Background())
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
