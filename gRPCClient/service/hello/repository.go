package hello

import "context"

type IgRPCHelloRepositoryClient interface {
	Hello(ctx context.Context, name string) (string, error)
}
