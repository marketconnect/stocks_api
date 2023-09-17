package card_data_provider

import (
	"context"
	pb "stocks_api/app/gen/proto"
	client "stocks_api/app/pkg/client/postgresql"
)

const (
	insertQuery = `INSERT INTO public.card (user_id, name, sku) VALUES ($1, $2, $3)`
	selectQuery = `SELECT * FROM public.card WHERE user_id = $1`
)

type cardStorage struct {
	client client.PostgreSQLClient
}

func NewCardStorage(client client.PostgreSQLClient) *cardStorage {
	return &cardStorage{client: client}
}

func (as *cardStorage) Save(ctx context.Context, userId uint64, cards []*pb.ProductCard) error {
	for _, card := range cards {
		name := card.GetName()
		sku := card.GetSku()

		_, err := as.client.Query(ctx, insertQuery, userId, name, sku)
		if err != nil {
			return err
		}
	}
	return nil
}
