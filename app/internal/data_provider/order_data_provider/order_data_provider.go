package order_data_provider

import (
	"context"
	pb "stocks_api/app/gen/proto"
	client "stocks_api/app/pkg/client/postgresql"
	"time"
)

const (
	selectQuery = `SELECT sku, qty FROM public.mc_order WHERE sku = ANY($1) AND created_at >= $2 AND created_at <= $3`
)

type orderStorage struct {
	client client.PostgreSQLClient
}

func NewOrderStorage(client client.PostgreSQLClient) *orderStorage {
	return &orderStorage{client: client}
}

func (as *orderStorage) GetOrdersFromTo(ctx context.Context, skus []uint64, dateFrom time.Time, dateTo time.Time) ([]*pb.Order, error) {
	var result []*pb.Order
	rows, err := as.client.Query(ctx, selectQuery, skus, dateFrom, dateTo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var order pb.Order
		// Scan the values from the row into the stock struct
		err := rows.Scan(&order.Sku, &order.Qty)
		if err != nil {
			return nil, err
		}
		result = append(result, &order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
