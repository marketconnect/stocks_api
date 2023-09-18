package auth_service

import (
	"context"

	"stocks_api/app/internal/domain/entity"
	my_jwt "stocks_api/app/pkg/jwt"
	"stocks_api/app/pkg/logger"

	pb "stocks_api/app/gen/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserStore interface {
	Save(ctx context.Context, user *entity.User) error
	Find(ctx context.Context, username string) (*entity.User, error)
}

type AuthService struct {
	store      UserStore
	logging    logger.Logger
	jwtManager my_jwt.JWTManager
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(store UserStore, jwtManager my_jwt.JWTManager, logging logger.Logger) *AuthService {
	return &AuthService{
		store:      store,
		jwtManager: jwtManager,
		logging:    logging,
	}
}

func (service *AuthService) Login(ctx context.Context, req *pb.AuthRequest) (*pb.TokenResponse, error) {
	user, err := service.store.Find(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := service.jwtManager.Generate(user.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pb.TokenResponse{Token: token}
	return res, nil
}

func (service *AuthService) Register(ctx context.Context, req *pb.AuthRequest) (*pb.TokenResponse, error) {
	username := req.GetUsername()
	pass := req.GetPassword()
	if (username == "") || (pass == "") {
		return nil, status.Errorf(codes.InvalidArgument, "username/password is empty")
	}

	newUser, err := entity.NewUser(username, pass)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create user: %v", err)
	}
	err = service.store.Save(ctx, newUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	token, err := service.jwtManager.Generate(newUser.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pb.TokenResponse{Token: token}
	return res, nil
}
