package main

import (
	"fmt"
	"github.com/jayson-hu/rpc/json_http/service"
	"io"
	"net/http"
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
func (h *HelloService) Hello(request string, response *string) error {
	*response = fmt.Sprintf("hello, %s", request)
	return nil

}


func NewRPCReadWriteCloserFromHTTP(w http.ResponseWriter, r *http.Request) *RPCReadWriteCloser {
	return &RPCReadWriteCloser{w, r.Body}
}

type RPCReadWriteCloser struct {
	io.Writer
	io.ReadCloser
}

func main() {
	rpc.RegisterName(service.SERVICE_NAME, new(HelloService))

	// RPC的服务架设在“/jsonrpc”路径，
	// 在处理函数中基于http.ResponseWriter和http.Request类型的参数构造一个io.ReadWriteCloser类型的conn通道。
	// 然后基于conn构建针对服务端的json编码解码器。
	// 最后通过rpc.ServeRequest函数为每次请求处理一次RPC方法调用
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		conn := NewRPCReadWriteCloserFromHTTP(w, r)
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":5060", nil)







}