package auth_data_provider

import (
	"context"

	"stocks_api/app/internal/domain/entity"
	client "stocks_api/app/pkg/client/postgresql"
)

const (
	saveQuery = `INSERT INTO public.mc_users (username, password) VALUES ($1, $2) RETURNING id`
	findQuery = `SELECT * FROM public.mc_users WHERE username = $1`
)

type authStorage struct {
	client client.PostgreSQLClient
}

func NewAuthStorage(client client.PostgreSQLClient) *authStorage {
	return &authStorage{client: client}
}

func (as *authStorage) Save(ctx context.Context, user *entity.User) error {

	row := as.client.QueryRow(ctx, saveQuery, user.Username, user.HashedPassword)
	var userID uint64
	err := row.Scan(&userID)

	return err
}

func (as *authStorage) Find(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	err := as.client.QueryRow(ctx, findQuery, username).Scan(&user)

	return &user, err
}
