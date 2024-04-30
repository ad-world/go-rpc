package rpcapi

import "fmt"

type ArithServer struct {}

func (t *ArithServer) Multiply(args *ArithArgs, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *ArithServer) Divide(args *ArithArgs, quo *Quotient) error {
	if args.B == 0 {
		return fmt.Errorf("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B

	return nil
}

func (t *ArithServer) Add(args *ArithArgs, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func (t *ArithServer) Sub(args *ArithArgs, reply *int) error {
	*reply = args.A - args.B
	return nil
}
