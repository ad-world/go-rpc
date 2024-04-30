package rpcapi

type Arith interface {
	Add(args *ArithArgs, reply *int) error
	Sub(args *ArithArgs, reply *int) error
	Multiply(args *ArithArgs, reply *int) error
	Divide(args *ArithArgs, quo *Quotient) error
}

type ArithArgs struct {
	A int
	B int
}

type Quotient struct {
	Quo int
	Rem int
}