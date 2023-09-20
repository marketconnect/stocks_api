package subscription_service

import (
	"context"
	"fmt"
	pb "stocks_api/app/gen/proto"
)

type SubscriptionStore interface {
	GetUserSubscriptionsByUserName(ctx context.Context, username string) ([]*pb.UserSubscription, error)
}

type SubscriptionService struct {
	store SubscriptionStore
	pb.UnimplementedUserSubscriptionsServiceServer
}

func NewSubscriptionService(store SubscriptionStore) *SubscriptionService {
	return &SubscriptionService{
		store: store,
	}
}

func (s *SubscriptionService) GetSubscriptions(ctx context.Context, req *pb.UserSubscriptionsRequest) (*pb.UserSubscriptionsResponse, error) {
	username := req.GetUsername()
	fmt.Println(username)
	subscriptions, err := s.store.GetUserSubscriptionsByUserName(ctx, username)
	if err != nil {
		return nil, err
	}
	return &pb.UserSubscriptionsResponse{
		Subscriptions: subscriptions,
	}, nil
}
