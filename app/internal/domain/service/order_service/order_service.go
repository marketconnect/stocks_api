package order_service

import (
	"context"
	pb "stocks_api/app/gen/proto"
	"stocks_api/app/pkg/logger"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderDataProvider interface {
	GetOrdersFromTo(ctx context.Context, skus []uint64, dateFrom time.Time, dateTo time.Time) ([]*pb.Order, error)
}

type OrderService struct {
	orderDataProvider OrderDataProvider
	logger            logger.Logger
	pb.UnimplementedOrderServiceServer
}

func NewOrderService(orderDataProvider OrderDataProvider, logger logger.Logger) *OrderService {
	return &OrderService{
		orderDataProvider: orderDataProvider,
		logger:            logger,
	}
}

func (service *OrderService) GetOrdersFromTo(ctx context.Context, req *pb.GetOrdersFromToReq) (*pb.GetOrdersFromToResp, error) {

	// Validate input parameters
	if req == nil {
		return &pb.GetOrdersFromToResp{}, status.Error(codes.InvalidArgument, "request is nil")
	}
	unixFrom := req.GetFrom()
	if unixFrom == 0 {
		return &pb.GetOrdersFromToResp{}, status.Error(codes.InvalidArgument, "from is required")
	}
	unixTo := req.GetTo()
	if unixTo == 0 {
		return &pb.GetOrdersFromToResp{}, status.Error(codes.InvalidArgument, "to is required")
	}
	skus := req.GetSkus()
	if len(skus) == 0 {
		return &pb.GetOrdersFromToResp{}, status.Error(codes.InvalidArgument, "skus is required")
	}
	from := time.Unix(0, unixFrom*1000000)
	to := time.Unix(0, unixTo*1000000)

	orders, err := service.orderDataProvider.GetOrdersFromTo(ctx, skus, from, to)
	if err != nil {
		service.logger.Error(err)
		return &pb.GetOrdersFromToResp{}, status.Errorf(codes.Internal, "could`t get orders: %v", err)
	}
	service.logger.Info(orders)
	return &pb.GetOrdersFromToResp{Orders: orders}, nil
}
