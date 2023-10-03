package auth_service

import (
	"context"
	"time"

	"stocks_api/app/internal/domain/entity"
	"stocks_api/app/pkg/logger"

	pb "stocks_api/app/gen/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserStore interface {
	Save(ctx context.Context, user *entity.User) (uint64, error)
	Find(ctx context.Context, username string) (*entity.User, error)
}

type TokenManager interface {
	Generate(userId uint64) (string, error)
}

type AuthService struct {
	userStore     UserStore
	tokenDuration int
	tokenManager  TokenManager
	logger        logger.Logger
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(userStore UserStore, tokenManager TokenManager, tokenDuration int, logger logger.Logger) *AuthService {
	return &AuthService{
		userStore: userStore,

		tokenManager:  tokenManager,
		tokenDuration: tokenDuration,
		logger:        logger,
	}
}

func (service *AuthService) Login(ctx context.Context, req *pb.AuthRequest) (*pb.TokenResponse, error) {
	userName := req.GetUsername()
	user, err := service.userStore.Find(ctx, userName)

	if err != nil {
		service.logger.Error(err)
		return nil, status.Errorf(codes.NotFound, "cannot find user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := service.tokenManager.Generate(user.Id)
	if err != nil {
		service.logger.Error(err)
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}
	apochNow := time.Now().Unix()

	res := &pb.TokenResponse{Token: token, ExpiredAt: apochNow + int64(service.tokenDuration)}
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
		service.logger.Error(err)
		return nil, status.Errorf(codes.InvalidArgument, "cannot create user: %v", err)
	}
	userId, err := service.userStore.Save(ctx, newUser)
	if err != nil {
		service.logger.Error(err)
		return nil, status.Errorf(codes.AlreadyExists, "cannot save user: %v", err)
	}

	token, err := service.tokenManager.Generate(userId)
	if err != nil {
		service.logger.Error(err)
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	apocheNow := time.Now().Unix()
	res := &pb.TokenResponse{Token: token, ExpiredAt: apocheNow + int64(service.tokenDuration)}
	return res, nil
}
