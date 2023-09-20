package subscription_data_provider

import (
	"context"
	"fmt"
	"time"

	pb "stocks_api/app/gen/proto"
	client "stocks_api/app/pkg/client/postgresql"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	selectByIdQuery   = `SELECT price, quantity, end_date FROM public.users_subscriptions WHERE user_id = $1 AND end_date > CURRENT_DATE`
	selectByNameQuery = `SELECT users_subscriptions.price, users_subscriptions.quantity, users_subscriptions.end_date FROM users_subscriptions JOIN mc_users ON users_subscriptions.user_id = mc_users.id WHERE mc_users.username = $1 AND end_date > CURRENT_DATE`
	insertQuery       = `INSERT INTO public.users_subscriptions (user_id, end_date, price, quantity) VALUES ($1, CURRENT_DATE + make_interval(days => $2), $3, $4)`
)

type subscriptionStorage struct {
	client client.PostgreSQLClient
}

func NewSubscriptionStorage(client client.PostgreSQLClient) *subscriptionStorage {
	return &subscriptionStorage{client: client}
}

func (s *subscriptionStorage) InsertSubscription(ctx context.Context, userId uint64, price float32, quantity, daysFromNow int32) error {

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

	if price < 0 {
		return fmt.Errorf("price must be greater than 0")
	}

	// Execute the query
	_ = s.client.QueryRow(ctx, insertQuery, userId, daysFromNow, price, quantity)

	return nil

}

func (s *subscriptionStorage) GetUserSubscriptionsByUserId(ctx context.Context, userId uint64) ([]*pb.UserSubscription, error) {

	rows, err := s.client.Query(ctx, selectByIdQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []*pb.UserSubscription
	for rows.Next() {
		var dateTo time.Time
		var qty int32
		var price float32
		err := rows.Scan(&price, &qty, &dateTo)
		if err != nil {
			return nil, err
		}

		timestampDateTo := timestamppb.Timestamp{
			Seconds: dateTo.Unix(),
			Nanos:   int32(dateTo.Nanosecond()),
		}

		subscription := &pb.UserSubscription{Price: price, Qty: qty, DateTo: &timestampDateTo}
		subscriptions = append(subscriptions, subscription)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func (s *subscriptionStorage) GetUserSubscriptionsByUserName(ctx context.Context, username string) ([]*pb.UserSubscription, error) {
	rows, err := s.client.Query(ctx, selectByNameQuery, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []*pb.UserSubscription
	for rows.Next() {
		var dateTo time.Time
		var qty int32
		var price float32
		err := rows.Scan(&price, &qty, &dateTo)
		if err != nil {
			return nil, err
		}

		timestampDateTo := timestamppb.Timestamp{
			Seconds: dateTo.Unix(),
			Nanos:   int32(dateTo.Nanosecond()),
		}
		subscription := &pb.UserSubscription{Price: price, Qty: qty, DateTo: &timestampDateTo}
		subscriptions = append(subscriptions, subscription)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return subscriptions, nil
}
