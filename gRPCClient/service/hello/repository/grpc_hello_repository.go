package repository

import (
	"client/proto/services"
	"client/service/hello"
	"context"
)

type gRPCHelloRepository struct {
	calculatorClient services.CalculatorClient
}

func NewgRPCHelloRepository(calculatorClient services.CalculatorClient) hello.IgRPCHelloRepositoryClient {
	return gRPCHelloRepository{
		calculatorClient: calculatorClient,
	}
}

func (g gRPCHelloRepository) Hello(ctx context.Context, name string) (string, error) {
	req := services.HelloRequest{
		Name: name,
	}

	resp, err := g.calculatorClient.Hello(ctx, &req)
	if err != nil {
		return "", err
	}

	return resp.Result, nil
}
