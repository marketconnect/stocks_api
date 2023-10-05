package card_data_provider

import (
	"context"
	"fmt"
	pb "stocks_api/app/gen/proto"
	client "stocks_api/app/pkg/client/postgresql"
)

const (
	insertQuery = `INSERT INTO public.card (user_id, name, sku, image) VALUES 	`
	selectQuery = `SELECT name, sku, image FROM public.card WHERE user_id = $1`
)

type cardStorage struct {
	client client.PostgreSQLClient
}

func NewCardStorage(client client.PostgreSQLClient) *cardStorage {
	return &cardStorage{client: client}
}

func (as *cardStorage) SaveAll(ctx context.Context, userId uint64, cards []*pb.ProductCard) (int32, error) {

	vals := []interface{}{}
	n := int32(0)
	i := 1
	sql := insertQuery
	for _, c := range cards {

		sql += fmt.Sprintf("($%d, $%d, $%d, $%d),", i, i+1, i+2, i+3)
		i = i + 4
		vals = append(vals, userId, c.Name, c.Sku, c.Image)
		n++
	}
	sql = sql[:len(sql)-1]

	sql += " ON CONFLICT (user_id, sku) DO UPDATE SET sku = EXCLUDED.sku"

	if _, err := as.client.Exec(ctx, sql, vals...); err != nil {
		return n, err
	}
	return n, nil

}

func (as *cardStorage) GetAll(ctx context.Context, userId uint64) ([]*pb.ProductCard, error) {
	var result []*pb.ProductCard

	rows, err := as.client.Query(ctx, selectQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var card pb.ProductCard

		// Scan the values from the row into the card struct
		err := rows.Scan(&card.Name, &card.Sku, &card.Image)
		if err != nil {
			// Handle the error here
			continue
		}

		result = append(result, &card)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
