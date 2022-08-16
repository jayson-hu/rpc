package main

import (
	"context"
	"fmt"
	"github.com/jayson-hu/rpc/grpc/simple/server/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

// HelloServiceServer is the server API for HelloService service.
//type HelloServiceServer interface {
//	Hello(context.Context, *Request) (*Response, error)
//}
type HelloServiceServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloServiceServer) Hello(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	return &pb.Response{
		Value: fmt.Sprintf("hello, %s", req.Value)}, nil
}
func main() {
	server := grpc.NewServer()
	//把实现类注册给grpc server
	//pb.RegisterHelloServiceServer(nil,&HelloServiceServer{})
	pb.RegisterHelloServiceServer(server,new(HelloServiceServer))
	//conn, err := net.Dial("tcp", ":1234")
	//if err != nil {
	//	panic(err)
	//}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	log.Printf("grpc listen addr:1234")
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}

}
