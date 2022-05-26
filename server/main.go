package main

import (
	"log"
	"math"
	"net"
	"net/rpc"
)

type Args struct {
	Number1, Number2 float64
}

type Calculator struct{}

func (c *Calculator) Sum(args *Args, result *float64) error {
	*result = args.Number1 + args.Number2

	return nil
}

func (c *Calculator) Sub(args *Args, result *float64) error {
	*result = args.Number1 - args.Number2

	return nil
}

func (c *Calculator) Mult(args *Args, result *float64) error {
	*result = args.Number1 * args.Number2

	return nil
}

func (c *Calculator) Div(args *Args, result *float64) error {
	*result = args.Number1 / args.Number2

	return nil
}

func (c *Calculator) Exp(args *Args, result *float64) error {
	*result = math.Pow(args.Number1, args.Number2)

	return nil
}

func main() {
	calculator := new(Calculator)

	rpc.Register(calculator)

	listener, e := net.Listen("tcp", ":3000")

	if e != nil {
		log.Fatalf("listen error: %v", e)
	}

	for {
		if conn, err := listener.Accept(); err != nil {
			log.Fatalf("accept error: %v", err.Error())
		} else {
			log.Printf("new connection establishedn")
			go rpc.ServeConn(conn)
		}
	}
}
