package permission_service

import (
	"context"
	pb "stocks_api/app/gen/proto"
)

type PermissionStore interface {
	GetUserPermissionsByUserName(ctx context.Context, username string) ([]*pb.UserPermission, error)
}

type PermissionService struct {
	store PermissionStore
	pb.UnimplementedUserPermissionsServiceServer
}

func NewPermissionService(store PermissionStore) *PermissionService {
	return &PermissionService{
		store: store,
	}
}

func (s *PermissionService) GetPermissions(ctx context.Context, req *pb.UserPermissionsRequest) (*pb.UserPermissionsResponse, error) {
	permissions, err := s.store.GetUserPermissionsByUserName(ctx, req.GetUsername())
	if err != nil {
		return nil, err
	}
	return &pb.UserPermissionsResponse{
		Permissions: permissions,
	}, nil
}
