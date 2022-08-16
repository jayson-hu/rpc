package main

import (
	"fmt"
	"github.com/jayson-hu/rpc/pbrpc/service"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//约束接口的实现
var _ service.HelloService = (*HelloService)(nil)
// service handler
type HelloService struct {

}

// Hello  方法 request 是请求  response 是响应
// request --> name
//response <-- hello name
func (h *HelloService) Hello(request  *service.Request, response *service.Response) error {
	response.Value = fmt.Sprintf("hello, %s", request)
	return nil

}

func main() {
	//RPC对外暴露的对象注册到rpc框架内部
	_ =  rpc.RegisterName(service.SERVICE_NAME, &HelloService{})


	//进行监听
	listener, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatal("Listen error TCP ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)

		}
		// server采用json来编解码
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}








}
