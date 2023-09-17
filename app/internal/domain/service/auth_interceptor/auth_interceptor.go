package auth_interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type AuthInterceptor struct{}

func NewAuthInterceptor() *AuthInterceptor {
	return &AuthInterceptor{}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		fmt.Println("--> unary interceptor: ", info.FullMethod)

		// TODO: implement authorization

		return handler(ctx, req)
	}
}
