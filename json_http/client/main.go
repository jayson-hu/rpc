package main

import (
	"fmt"
	"github.com/jayson-hu/rpc/json_http/service"
	"log"
	"net/rpc"
)

//约束客户端
var _ service.HelloService = (*HelloServiceClient)(nil)

func NewHelloServiceClient(network, address string) (*HelloServiceClient, error ){
	client, err := rpc.Dial(network, address)
	if err != nil {
		log.Fatal("client fatal")
		return nil, err
	}
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


}
