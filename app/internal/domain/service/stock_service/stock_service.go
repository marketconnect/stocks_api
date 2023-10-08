package stock_service

import (
	"context"
	"fmt"
	pb "stocks_api/app/gen/proto"
	"stocks_api/app/pkg/logger"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StockDataProvider interface {
	GetStocksFromTo(ctx context.Context, skus []uint64, dateFrom time.Time, dateTo time.Time) ([]*pb.Stock, error)
}

type StockService struct {
	stockDataProvider StockDataProvider
	logger            logger.Logger
	pb.UnimplementedStockServiceServer
}

func NewStockService(stockDataProvider StockDataProvider, logger logger.Logger) *StockService {
	return &StockService{
		stockDataProvider: stockDataProvider,
		logger:            logger,
	}
}

func (service *StockService) GetStocksFromTo(ctx context.Context, req *pb.GetStocksFromToReq) (*pb.GetStocksFromToResp, error) {
	// Validate input parameters
	if req == nil {
		return &pb.GetStocksFromToResp{}, status.Error(codes.InvalidArgument, "request is nil")
	}
	unixFrom := req.GetFrom()
	if unixFrom == 0 {
		return &pb.GetStocksFromToResp{}, status.Error(codes.InvalidArgument, "from is required")
	}
	unixTo := req.GetTo()
	if unixTo == 0 {
		return &pb.GetStocksFromToResp{}, status.Error(codes.InvalidArgument, "to is required")
	}
	skus := req.GetSkus()
	if len(skus) == 0 {
		return &pb.GetStocksFromToResp{}, status.Error(codes.InvalidArgument, "skus is required")
	}

	from := time.Unix(int64(unixFrom), 0)
	to := time.Unix(int64(unixTo), 0)

	fmt.Printf("from: %v, to: %v\n", from, to)
	stocks, err := service.stockDataProvider.GetStocksFromTo(ctx, skus, from, to)
	if err != nil {
		service.logger.Error(err)
		return &pb.GetStocksFromToResp{}, status.Errorf(codes.Internal, "could`t get stocks: %v", err)
	}
	return &pb.GetStocksFromToResp{Stocks: stocks}, nil
}
