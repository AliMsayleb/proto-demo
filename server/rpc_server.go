package server

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sef-demo/rpc"
	"sef-demo/services"
)

func StartGRPC() {
	server := grpc.NewServer()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	rpc.RegisterUserServiceServer(server, services.UserRPCService{})
	rpc.RegisterCalculatorServiceServer(server, services.UserRPCService{})
	rpc.RegisterPizzaServiceServer(server, services.PizzaRPCService{})
	go func() {
		serveErr := server.Serve(listener)
		if serveErr != nil {
			fmt.Printf("Couldn't shut down GRPC server properly %s\n", serveErr)
		}
	}()
}
