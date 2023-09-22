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
	GetAllUserSubscriptionsByUserId(ctx context.Context, userId uint64) ([]*pb.UserSubscription, error)
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
	userId := req.GetID()
	fmt.Printf("HERE %d", userId)
	if userId == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "userID is zero")
	}
	fmt.Println(userId)
	subscriptions, err := s.store.GetAllUserSubscriptionsByUserId(ctx, userId)
	if err != nil {
		s.logger.Error(err)
		return nil, status.Errorf(codes.Internal, "cannot find subsciptions: %v", err)
	}
	return &pb.UserSubscriptionsResponse{
		Subscriptions: subscriptions,
	}, nil
}
