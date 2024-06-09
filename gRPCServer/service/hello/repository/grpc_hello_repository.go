package repository

import (
	"context"
	"fmt"
	"server/proto/services"
	"server/service/hello"
)

type gRPCHelloRepository struct{}

func NewgRPCHelloRepository() hello.IgRPCHelloRepository {
	return gRPCHelloRepository{}
}

func (g gRPCHelloRepository) Hello(ctx context.Context, req *services.HelloRequest) (*services.HelloResponse, error) {
	result := fmt.Sprintf("Hello %s", req.Name)
	resp := &services.HelloResponse{
		Result: result,
	}
	return resp, nil
}
