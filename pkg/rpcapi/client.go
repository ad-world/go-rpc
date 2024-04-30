package rpcapi

import "net/rpc"

type ArithClient struct {
	*rpc.Client
}

func (ac *ArithClient) Add(args *ArithArgs, reply *int) error {
	return ac.Call("ArithServer.Add", args, reply)
}

func (ac *ArithClient) Sub(args *ArithArgs, reply *int) error {
	return ac.Call("ArithServer.Sub", args, reply)
}

func (ac *ArithClient) Multiply(args *ArithArgs, reply *int) error {
	return ac.Call("ArithServer.Multiply", args, reply)
}

func (ac *ArithClient) Divide(args *ArithArgs, quo *Quotient) error {
	return ac.Call("ArithServer.Divide", args, quo)
}



