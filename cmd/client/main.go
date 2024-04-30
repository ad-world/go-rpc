package main

import (
	"fmt"
	"go-rpc/pkg/rpcapi"
	"net/rpc"
)

func main() {
	c, _ := rpc.Dial("tcp", "localhost:1234")
	client := &rpcapi.ArithClient{Client: c}

	args := &rpcapi.ArithArgs{A: 7, B: 8}
	var add_reply int
	var mult_reply int
	var sub_reply int
	var quo_reply rpcapi.Quotient

	client.Add(args, &add_reply)
	fmt.Println("Add: ", add_reply)
	client.Divide(args, &quo_reply)
	fmt.Println("Divide: ", quo_reply)
	client.Sub(args, &sub_reply)
	fmt.Println("Sub: ", sub_reply)
	client.Multiply(args, &mult_reply)
	fmt.Println("Multiply: ", mult_reply)
}