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
	selectActiveByIdQuery = `SELECT quantity, price, end_date FROM public.users_subscriptions WHERE user_id = $1 AND end_date > CURRENT_DATE`
	selectAllByNameQuery  = `SELECT price, quantity, end_date, info, created_at FROM users_subscriptions WHERE user_id = $1`
	insertQuery           = `INSERT INTO public.users_subscriptions (user_id, end_date, price, info, quantity) VALUES ($1, CURRENT_DATE + make_interval(days => $2), $3, $4, $5)`
)

type subscriptionStorage struct {
	client client.PostgreSQLClient
}

func NewSubscriptionStorage(client client.PostgreSQLClient) *subscriptionStorage {
	return &subscriptionStorage{client: client}
}

func (s *subscriptionStorage) InsertSubscription(ctx context.Context, userId uint64, price float32, info string, quantity, daysFromNow int32) error {

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

	if info == "" {
		return fmt.Errorf("info must be not empty")
	}

	// Execute the query
	_ = s.client.QueryRow(ctx, insertQuery, userId, daysFromNow, price, info, quantity)

	return nil

}

func (s *subscriptionStorage) GetActiveUserSubscriptionsByUserId(ctx context.Context, userId uint64) ([]*pb.UserSubscription, error) {

	fmt.Println(userId)
	rows, err := s.client.Query(ctx, selectActiveByIdQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []*pb.UserSubscription
	for rows.Next() {
		var dateTo time.Time
		var qty int32
		var price float32
		err := rows.Scan(&qty, &price, &dateTo)
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

func (s *subscriptionStorage) GetAllUserSubscriptionsByUserId(ctx context.Context, userId uint64) ([]*pb.UserSubscription, error) {
	rows, err := s.client.Query(ctx, selectAllByNameQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []*pb.UserSubscription
	for rows.Next() {
		var qty int32
		var price float32
		var dateTo time.Time
		var info string
		var createdAt time.Time
		err := rows.Scan(&price, &qty, &dateTo, &info, &createdAt)
		if err != nil {
			return nil, err
		}

		timestampDateTo := timestamppb.Timestamp{
			Seconds: dateTo.Unix(),
			Nanos:   int32(dateTo.Nanosecond()),
		}

		timestampCreatedAt := timestamppb.Timestamp{
			Seconds: createdAt.Unix(),
			Nanos:   int32(createdAt.Nanosecond()),
		}

		subscription := &pb.UserSubscription{Price: price, Qty: qty, DateTo: &timestampDateTo, Info: info, CreatedAt: &timestampCreatedAt}
		subscriptions = append(subscriptions, subscription)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return subscriptions, nil
}
