package main

import (
	"client/proto/services"
	_hello_repository "client/service/hello/repository"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	credent := insecure.NewCredentials()
	gRPCClient, err := grpc.NewClient("localhost:3100", grpc.WithTransportCredentials(credent))
	if err != nil {
		panic(err)
	}
	defer gRPCClient.Close()

	helloRepo := _hello_repository.NewgRPCHelloRepository(services.NewCalculatorClient(gRPCClient))
	result, err := helloRepo.Hello(ctx, "Pheet")
	if err != nil {
		panic(err)
	}

	fmt.Println("Result:", result)
}
