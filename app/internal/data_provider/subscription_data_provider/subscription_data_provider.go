package subscription_data_provider

import (
	"context"
	"fmt"

	pb "stocks_api/app/gen/proto"
	client "stocks_api/app/pkg/client/postgresql"
)

const (
	selectByIdQuery   = `SELECT quantity, end_date FROM public.users_subscriptions WHERE user_id = $1 AND end_date > CURRENT_DATE`
	selectByNameQuery = `SELECT users_subscription.quantity, users_subscription.end_date FROM users_subscription JOIN mc_users ON users_subscription.user_id = mc_users.id WHERE mc_users.username = $1 AND end_date > CURRENT_DATE`
	insertQuery       = `INSERT INTO public.users_subscriptions (user_id, end_date, quantity) VALUES ($1, CURRENT_DATE + make_interval(days => $2), $3)`
)

type subscriptionStorage struct {
	client client.PostgreSQLClient
}

func NewSubscriptionStorage(client client.PostgreSQLClient) *subscriptionStorage {
	return &subscriptionStorage{client: client}
}

func (s *subscriptionStorage) InsertSubscription(ctx context.Context, userId uint64, quantity, daysFromNow int32) error {

	// Validate arguments
	if userId == 0 {
		return fmt.Errorf("userId must be greater than 0")
	}
	if quantity <= 0 {
		return fmt.Errorf("quantity must be greater than 0")
	}
	if daysFromNow <= 0 {
		return fmt.Errorf("daysFromNow must be greater than 0")
	}

	// Execute the query
	_ = s.client.QueryRow(ctx, insertQuery, userId, daysFromNow, quantity)

	return nil

}

func (s *subscriptionStorage) GetUserSubscriptionByUserId(ctx context.Context, userId uint64) ([]*pb.UserSubscription, error) {

	rows, err := s.client.Query(ctx, selectByIdQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []*pb.UserSubscription
	for rows.Next() {
		subscription := &pb.UserSubscription{}
		err := rows.Scan(&subscription.Qty, &subscription.DateTo)
		if err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, subscription)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func (s *subscriptionStorage) GetUserSubscriptionByUserName(ctx context.Context, username string) ([]*pb.UserSubscription, error) {
	rows, err := s.client.Query(ctx, selectByNameQuery, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []*pb.UserSubscription
	for rows.Next() {
		subscription := &pb.UserSubscription{}
		err := rows.Scan(&subscription.Qty, &subscription.DateTo)
		if err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, subscription)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return subscriptions, nil
}
