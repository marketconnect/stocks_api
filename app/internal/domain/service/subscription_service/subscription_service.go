package subscription_service

import (
	"context"
	pb "stocks_api/app/gen/proto"
)

type SubscriptionStore interface {
	GetUserSubscriptionByUserName(ctx context.Context, username string) ([]*pb.UserSubscription, error)
}

type SubscriptionService struct {
	store SubscriptionStore
	pb.UnimplementedUserSubscriptionServiceServer
}

func NewSubscriptionService(store SubscriptionStore) *SubscriptionService {
	return &SubscriptionService{
		store: store,
	}
}

func (s *SubscriptionService) GetSubscription(ctx context.Context, req *pb.UserSubscriptionRequest) (*pb.UserSubscriptionResponse, error) {
	subscriptions, err := s.store.GetUserSubscriptionByUserName(ctx, req.GetUsername())
	if err != nil {
		return nil, err
	}
	return &pb.UserSubscriptionResponse{
		Subscription: subscriptions,
	}, nil
}
