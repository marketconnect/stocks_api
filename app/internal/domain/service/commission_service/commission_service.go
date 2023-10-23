package commission_service

import (
	"context"
	pb "stocks_api/app/gen/proto"
	"stocks_api/app/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CommissionDataProvider interface {
	GetCommission(ctx context.Context, id uint64) (*pb.GetCommissionResp, error)
}

type CommissionService struct {
	commissionDataProvider CommissionDataProvider
	logger                 logger.Logger
	pb.UnimplementedCommissionServiceServer
}

func NewCommissionService(commissionDataProvider CommissionDataProvider, logger logger.Logger) *CommissionService {
	return &CommissionService{
		commissionDataProvider: commissionDataProvider,
		logger:                 logger,
	}
}

func (service *CommissionService) GetStocksFromTo(ctx context.Context, req *pb.GetCommissionReq) (*pb.GetCommissionResp, error) {
	// Validate input parameters
	if req == nil {
		return &pb.GetCommissionResp{}, status.Error(codes.InvalidArgument, "request is nil")
	}
	id := req.GetId()
	if id == 0 {
		return &pb.GetCommissionResp{}, status.Error(codes.InvalidArgument, "id is required")
	}

	commission, err := service.commissionDataProvider.GetCommission(ctx, id)
	if err != nil {
		service.logger.Error(err)
		return &pb.GetCommissionResp{}, status.Errorf(codes.Internal, "could`t get commission: %v", err)
	}
	return commission, nil
}
