package subscription_service

import (
	"context"
	"fmt"
	pb "stocks_api/app/gen/proto"
	"stocks_api/app/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SubscriptionStore interface {
	GetUserSubscriptionsByUserName(ctx context.Context, username string) ([]*pb.UserSubscription, error)
}

type SubscriptionService struct {
	store  SubscriptionStore
	logger logger.Logger
	pb.UnimplementedUserSubscriptionsServiceServer
}

func NewSubscriptionService(store SubscriptionStore, logger logger.Logger) *SubscriptionService {
	return &SubscriptionService{
		store:  store,
		logger: logger,
	}
}

func (s *SubscriptionService) GetSubscriptions(ctx context.Context, req *pb.UserSubscriptionsRequest) (*pb.UserSubscriptionsResponse, error) {
	username := req.GetUsername()
	fmt.Println("HERE " + username)
	if username == "" {
		return nil, status.Errorf(codes.InvalidArgument, "username is empty")
	}
	fmt.Println(username)
	subscriptions, err := s.store.GetUserSubscriptionsByUserName(ctx, username)
	if err != nil {
		s.logger.Error(err)
		return nil, status.Errorf(codes.Internal, "cannot find subsciptions: %v", err)
	}
	return &pb.UserSubscriptionsResponse{
		Subscriptions: subscriptions,
	}, nil
}
