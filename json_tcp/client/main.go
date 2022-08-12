package main

import (
	"fmt"
	"github.com/jayson-hu/rpc/json_tcp/service"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//约束客户端
var _ service.HelloService = (*HelloServiceClient)(nil)

func NewHelloServiceClient(network, address string) (*HelloServiceClient, error ){
	conn, err := net.Dial(network, address)
	if err != nil {
		log.Fatal("client fatal")
		return nil, err
	}
	// 客户端采用json格式来编解码
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	return &HelloServiceClient{
		client:client,
	}, nil
}


type HelloServiceClient struct {
	client *rpc.Client
}

func (c *HelloServiceClient) Hello(request string, response *string) error {
	//var resp string
	//err = c.client.Call("HelloService.Hello", "alice1", &resp)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(resp)
	//return nil
	return c.client.Call(fmt.Sprintf("%s.Hello",service.SERVICE_NAME), request, response)
}

func main() {
	c, err := NewHelloServiceClient("tcp", "localhost:5060")
	if err != nil {
		return
	}
	var resp string
	err = c.Hello("bol", &resp)
	if err != nil {
		return
	}
	fmt.Println(resp)

}
