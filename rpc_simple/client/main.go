package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:5060")
	if err != nil {
		log.Fatal("client")
	}
	var resp string
	err = client.Call("HelloService.Hello", "alice1", &resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
