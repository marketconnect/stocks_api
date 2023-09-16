package auth_service

import (
	"context"
	"fmt"

	my_jwt "stocks_api/app/internal/domain/jwt"
	"stocks_api/app/pkg/logger"

	pb "stocks_api/app/gen/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthStorage interface {
	RegisterUser(ctx context.Context, email, password string) (uint64, error)
	LoginUser(ctx context.Context, email, password string) (uint64, error)
}

type AuthService struct {
	storage AuthStorage
	logging logger.Logger
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(storage AuthStorage, logging logger.Logger) *AuthService {
	return &AuthService{
		storage: storage,
		logging: logging,
	}
}

func (s *AuthService) RegisterUser(ctx context.Context, user *pb.User) (*pb.TokenMessage, error) {
	email := user.GetEmail()
	pswd := user.GetPassword()
	id, err := s.storage.RegisterUser(ctx, email, pswd)
	fmt.Println(email, pswd)
	if err != nil {
		return nil, err
	}
	token, err := my_jwt.CreateToken(id)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")

	}

	return &pb.TokenMessage{Token: token}, nil
}

func (s *AuthService) LoginUser(ctx context.Context, user *pb.User) (*pb.TokenMessage, error) {
	email := user.GetEmail()
	pswd := user.GetPassword()
	id, err := s.storage.LoginUser(ctx, email, pswd)
	if err != nil {

		return nil, err
	}
	token, err := my_jwt.CreateToken(id)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &pb.TokenMessage{Token: token}, nil
}
