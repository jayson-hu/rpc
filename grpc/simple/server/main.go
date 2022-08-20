package main

import (
	"context"
	"fmt"
	"github.com/jayson-hu/rpc/grpc/middleware/server"
	"github.com/jayson-hu/rpc/grpc/simple/server/pb"
	"google.golang.org/grpc"
	"io"
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

func (s *HelloServiceServer) Channel(req pb.HelloService_ChannelServer) error {
	for  {
		//接受请求
		recv, err := req.Recv()
		if err != nil {
			log.Printf("recv error, %s", err)
			if err == io.EOF{
				log.Printf("recv cloeds %s",err)

				return nil
			}
			return err
		}
		resp := &pb.Response{
			Value: fmt.Sprintf("hello, %s", recv.Value),
		}

		//相应请求
		err = req.Send(resp)
		if err != nil {
			if err == io.EOF{
				log.Printf("client closed err :",err)
				return nil
			}
			return err
		}

	}
}




func main() {
	reqauth := server.NewAuthUnaryServerInterceptor()
	streamAuth := server.NewAuthStreamUnaryServerInterceptor()
	server := grpc.NewServer(grpc.UnaryInterceptor(reqauth),grpc.StreamInterceptor(streamAuth))
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
