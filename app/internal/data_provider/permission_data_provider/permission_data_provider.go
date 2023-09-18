package permission_data_provider

import (
	"context"

	pb "stocks_api/app/gen/proto"
	client "stocks_api/app/pkg/client/postgresql"
)

const (
	selectByIdQuery   = `SELECT method, qty, date_to FROM user_permissions WHERE user_id = $1`
	selectByNameQuery = `SELECT user_permissions.* FROM user_permissions JOIN mc_users ON user_permissions.user_id = mc_users.id WHERE mc_users.username = $1`
)

type permissionStorage struct {
	client client.PostgreSQLClient
}

func NewPermissionStorage(client client.PostgreSQLClient) *permissionStorage {
	return &permissionStorage{client: client}
}

func (s *permissionStorage) GetUserPermissionsByUserId(ctx context.Context, userId uint64) ([]*pb.UserPermission, error) {
	rows, err := s.client.Query(ctx, selectByIdQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*pb.UserPermission
	for rows.Next() {
		permission := &pb.UserPermission{}
		err := rows.Scan(&permission.Method, &permission.Qty, &permission.DateTo)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (s *permissionStorage) GetUserPermissionsByUserName(ctx context.Context, username string) ([]*pb.UserPermission, error) {
	rows, err := s.client.Query(ctx, selectByIdQuery, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*pb.UserPermission
	for rows.Next() {
		permission := &pb.UserPermission{}
		err := rows.Scan(&permission.Method, &permission.Qty, &permission.DateTo)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}
