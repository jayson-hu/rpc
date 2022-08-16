package main

import (
	"context"
	"fmt"
	"github.com/jayson-hu/rpc/grpc/simple/server/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial( "localhost:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	//grpc生成一个client
	client :=pb.NewHelloServiceClient(conn)
	response, err := client.Hello(context.Background(),&pb.Request{
		Value: "alice",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Value)
}
