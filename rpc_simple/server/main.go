package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// service handler
type HelloService struct {
	
}

// Hello  方法 request 是请求  response 是响应
// request --> name
//response <-- hello name
func (h *HelloService) Hello(request string, response *string) error {
	*response = fmt.Sprintf("hello, %s", request)
	return nil

}

func main() {
	//RPC对外暴露的对象注册到rpc框架内部
	_ =  rpc.RegisterName("HelloService", &HelloService{})


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
		go rpc.ServeConn(conn)
	}








}