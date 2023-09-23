package stock_data_provider

import (
	"context"
	pb "stocks_api/app/gen/proto"
	client "stocks_api/app/pkg/client/postgresql"
	"time"
)

const (
	selectQuery = `SELECT sku, wh, qty FROM public.stock WHERE sku = $1 AND created_at >= $2 AND created_at <= $3`
)

type stockStorage struct {
	client client.PostgreSQLClient
}

func NewStockStorage(client client.PostgreSQLClient) *stockStorage {
	return &stockStorage{client: client}
}

func (as *stockStorage) GetFromTo(ctx context.Context, skus []uint64, dateFrom time.Time, dateTo time.Time) ([]*pb.Stock, error) {
	var result []*pb.Stock
	rows, err := as.client.Query(ctx, selectQuery, skus, dateFrom, dateTo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var stock pb.Stock
		// Scan the values from the row into the stock struct
		err := rows.Scan(&stock.Sku, &stock.Wh, &stock.Qty)
		if err != nil {
			return nil, err
		}
		result = append(result, &stock)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
