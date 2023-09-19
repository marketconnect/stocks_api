package auth_service

import (
	"context"

	"stocks_api/app/internal/domain/entity"

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
	store        UserStore
	tokenManager TokenManager
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(store UserStore, tokenManager TokenManager) *AuthService {
	return &AuthService{
		store:        store,
		tokenManager: tokenManager,
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

	token, err := service.tokenManager.Generate(user.Id)
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
		return nil, status.Errorf(codes.InvalidArgument, "cannot create user: %v", err)
	}
	userId, err := service.store.Save(ctx, newUser)
	if err != nil {
		return nil, status.Errorf(codes.AlreadyExists, "cannot find user: %v", err)
	}

	token, err := service.tokenManager.Generate(userId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pb.TokenResponse{Token: token}
	return res, nil
}
