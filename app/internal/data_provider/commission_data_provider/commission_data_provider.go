package commission_data_provider

import (
	"context"
	pb "stocks_api/app/gen/proto"
	client "stocks_api/app/pkg/client/postgresql"
)

const (
	selectQuery = `select category, subject, commission, fbs, fbo from commission where id =  $1`
)

type commissionStorage struct {
	client client.PostgreSQLClient
}

func NewCommissionStorage(client client.PostgreSQLClient) *commissionStorage {
	return &commissionStorage{client: client}
}

func (cs *commissionStorage) GetCommission(ctx context.Context, id uint64) (*pb.GetCommissionResp, error) {
	var result pb.GetCommissionResp
	row := cs.client.QueryRow(ctx, selectQuery, id)
	err := row.Scan(&result.Category, &result.Subject, &result.Commission, &result.Fbs, &result.Fbo)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
