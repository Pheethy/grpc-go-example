package main

import (
	"fmt"
	"net"
	"server/proto/services"
	_hello_repository "server/service/hello/repository"

	"google.golang.org/grpc"
)

const (
	GRPC_PORT = "3100"
)

func main() {
	gRPCServer := grpc.NewServer()

	hellRepo := _hello_repository.NewgRPCHelloRepository()

	services.RegisterCalculatorServer(gRPCServer, hellRepo)

	startGRPCServer(gRPCServer)
}

func startGRPCServer(server *grpc.Server) {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", GRPC_PORT))
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	/* serve grpc */
	fmt.Println(fmt.Sprintf("Start grpc Server [::%s]", GRPC_PORT))
	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
