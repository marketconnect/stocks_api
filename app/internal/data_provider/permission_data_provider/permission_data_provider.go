package permission_data_provider

import (
	"context"

	"stocks_api/app/internal/domain/entity"
	client "stocks_api/app/pkg/client/postgresql"
)

const (
	selectQuery = `SELECT user_permissions.* FROM user_permissions JOIN mc_users ON user_permissions.usert_id = mc_users.id	WHERE mc_users.username = $1`
)

type permissionStorage struct {
	client client.PostgreSQLClient
}

func NewPermissionStorage(client client.PostgreSQLClient) *permissionStorage {
	return &permissionStorage{client: client}
}

func (s *permissionStorage) GetUserPermissionsByUserName(ctx context.Context, userName string) ([]*entity.UserPermission, error) {
	rows, err := s.client.Query(ctx, selectQuery, userName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*entity.UserPermission
	for rows.Next() {
		permission := &entity.UserPermission{}
		err := rows.Scan(&permission.ID, &permission.Method, &permission.Qty, &permission.DateTo, &permission.UserID)
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
