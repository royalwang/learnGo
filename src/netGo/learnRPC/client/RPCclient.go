package main

import (
	"fmt"
	"net/rpc"
	"netGo/learnRPC/server"
	"reflect"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("err")
	}
	BlockRPC(client)
	NoneBlockRPC(client)
}

func BlockRPC(c *rpc.Client) {

	args := &server.Args{7, 8}
	var reply int
	err := c.Call("Arith.Multiply", args, &reply)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}

func NoneBlockRPC(c *rpc.Client) {
	args := &server.Args{7, 8}
	quotient := new(server.Quotient)
	divCall := c.Go("Arith.Divide", args, &quotient, nil)
	fmt.Println(quotient.Quo)
	fmt.Println(quotient.Rem)
	replyCall := <-divCall.Done
	fmt.Println(quotient.Quo)
	fmt.Println(quotient.Rem)

//	fmt.Println(reflect.TypeOf(replyCall))
	fmt.Println(replyCall.ServiceMethod)
	fmt.Println(reflect.TypeOf(replyCall.Reply))
	fmt.Println(replyCall.Reply)
//	fmt.Println(quotient.Quo)
//	fmt.Println(quotient.Rem)
}
