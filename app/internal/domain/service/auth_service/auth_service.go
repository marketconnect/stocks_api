package auth_service

import (
	"context"

	"stocks_api/app/internal/domain/entity"
	"stocks_api/app/pkg/logger"

	pb "stocks_api/app/gen/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SubscriptionStore interface {
	InsertSubscription(ctx context.Context, userId uint64, price float32, info string, quantity, daysFromNow int32) error
}

type UserStore interface {
	Save(ctx context.Context, user *entity.User) (uint64, error)
	Find(ctx context.Context, username string) (*entity.User, error)
}

type TokenManager interface {
	Generate(userId uint64) (string, error)
}

type AuthService struct {
	userStore         UserStore
	subscriptionStore SubscriptionStore
	tokenManager      TokenManager
	logger            logger.Logger
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(userStore UserStore, subscriptionStore SubscriptionStore, tokenManager TokenManager, logger logger.Logger) *AuthService {
	return &AuthService{
		userStore:         userStore,
		subscriptionStore: subscriptionStore,
		tokenManager:      tokenManager,
		logger:            logger,
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

	res := &pb.TokenResponse{Token: token, UserId: user.Id}
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

	err = service.subscriptionStore.InsertSubscription(ctx, userId, 0, "Бесплатный пробный период 7 дней", 10, 7)
	if err != nil {
		service.logger.Error(err)
		return nil, status.Errorf(codes.Internal, "cannot generate subscription")
	}

	res := &pb.TokenResponse{Token: token, UserId: userId}
	return res, nil
}
