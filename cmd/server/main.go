package main

import (
	"go-rpc/pkg/rpcapi"
	"net"
	"net/rpc"
)

func main() {
	server := new(rpcapi.ArithServer)

	rpc.Register(server)

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	defer l.Close()

	rpc.Accept(l)
}