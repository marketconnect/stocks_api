package auth_interceptor

import (
	"context"
	pb "stocks_api/app/gen/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type PermissionsStore interface {
	GetUserPermissionsByUserId(ctx context.Context, userId uint64) ([]*pb.UserPermission, error)
}

type TokenManager interface {
	Verify(accessToken string) (*uint64, error)
}

type AuthInterceptor struct {
	permissionsStore PermissionsStore
	tokenManager     TokenManager
}

func NewAuthInterceptor(permissionsStore PermissionsStore, tokenManager TokenManager) *AuthInterceptor {
	return &AuthInterceptor{permissionsStore: permissionsStore, tokenManager: tokenManager}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	userId, err := interceptor.tokenManager.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	userPermissions, err := interceptor.permissionsStore.GetUserPermissionsByUserId(ctx, *userId)
	if err != nil {
		return status.Error(codes.PermissionDenied, "no permission to access this RPC: "+method)
	}

	nowTime := time.Now()
	for _, permission := range userPermissions {
		if permission.Method == method && permission.DateTo.AsTime().After(nowTime) {
			return nil
		}
	}

	return status.Error(codes.PermissionDenied, "no permission to access this RPC: "+method)

}
