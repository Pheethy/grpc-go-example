package hello

import (
	"context"
	"server/proto/services"
)

type IgRPCHelloRepository interface {
	Hello(context.Context, *services.HelloRequest) (*services.HelloResponse, error)
}
