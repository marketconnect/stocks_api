package auth_interceptor

import (
	"context"
	"stocks_api/app/internal/domain/entity"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type PermissionsStore interface {
	GetUserPermissionsByUserName(ctx context.Context, userName string) ([]*entity.UserPermission, error)
}

type TokenManager interface {
	Verify(accessToken string) (*string, error)
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
	userName, err := interceptor.tokenManager.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	// accessibleRoles, ok := interceptor.permissionsStore.GetUserPermissionsByUserName(ctx, userName)
	userPermissions, err := interceptor.permissionsStore.GetUserPermissionsByUserName(ctx, *userName)
	if err != nil {
		return status.Error(codes.PermissionDenied, "no permission to access this RPC")
	}
	// Check if method accessible for the user
	nowTime := time.Now()
	for _, permission := range userPermissions {
		if permission.Method == method && permission.DateTo.After(nowTime) {
			return nil
		}
	}

	return status.Error(codes.PermissionDenied, "no permission to access this RPC")

}
