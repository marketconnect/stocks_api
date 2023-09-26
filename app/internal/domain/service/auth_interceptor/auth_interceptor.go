package auth_interceptor

import (
	"context"
	"fmt"
	pb "stocks_api/app/gen/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type SubscriptionStore interface {
	GetActiveUserSubscriptionsByUserId(ctx context.Context, userId uint64) ([]*pb.UserSubscription, error)
}

type TokenManager interface {
	Verify(accessToken string) (*uint64, error)
}

type AuthInterceptor struct {
	subscriptionStore SubscriptionStore
	tokenManager      TokenManager
}

func NewAuthInterceptor(subscriptionStore SubscriptionStore, tokenManager TokenManager) *AuthInterceptor {
	return &AuthInterceptor{subscriptionStore: subscriptionStore, tokenManager: tokenManager}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		fmt.Println(info.FullMethod)
		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	// Do not authorize for registration
	if method == "/main.AuthService/Register" || method == "/main.AuthService/Login" || method == "/main.UserSubscriptionsService/GetSubscriptions" {
		return nil
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		fmt.Println(0000000)
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided"+method)
	}

	accessToken := values[0]
	userId, err := interceptor.tokenManager.Verify(accessToken)
	if err != nil {
		fmt.Println(err)
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	userPermissions, err := interceptor.subscriptionStore.GetActiveUserSubscriptionsByUserId(ctx, *userId)
	if err != nil {
		fmt.Println(err)
		return status.Error(codes.PermissionDenied, "error to  permission to access this RPC: "+err.Error()+method)
	}

	if len(userPermissions) > 0 {
		fmt.Println("good")
		return nil
	}

	return status.Error(codes.PermissionDenied, "no permission to access this RPC: "+method)

}
