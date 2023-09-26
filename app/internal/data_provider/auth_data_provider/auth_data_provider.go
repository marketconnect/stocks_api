package auth_data_provider

import (
	"context"
	"errors"
	"fmt"

	"stocks_api/app/internal/domain/entity"
	client "stocks_api/app/pkg/client/postgresql"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

const (
	saveQuery = `INSERT INTO public.mc_users (username, password) VALUES ($1, $2) ON CONFLICT (username) DO UPDATE SET username = EXCLUDED.username	RETURNING id;
	`
	findQuery = `SELECT id, password FROM public.mc_users WHERE username = $1`
)

type authStorage struct {
	client client.PostgreSQLClient
}

func NewAuthStorage(client client.PostgreSQLClient) *authStorage {
	return &authStorage{client: client}
}

func (as *authStorage) Save(ctx context.Context, user *entity.User) (uint64, error) {
	row := as.client.QueryRow(ctx, saveQuery, user.Username, user.Password)
	var userID uint64
	err := row.Scan(&userID)

	if err != nil {
		var e *pgconn.PgError
		if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
			return userID, nil
		}

		return 0, fmt.Errorf("cannot save user: %d %w", user.Id, err)
	}

	return userID, nil
}
func (as *authStorage) Find(ctx context.Context, username string) (*entity.User, error) {
	var id uint64
	var password string
	err := as.client.QueryRow(ctx, findQuery, username).Scan(&id, &password)
	user := entity.User{Id: id, Username: username, Password: password}
	return &user, err
}
