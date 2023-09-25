package card_service

import (
	"context"
	"fmt"
	pb "stocks_api/app/gen/proto"
	"stocks_api/app/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CardDataProvider interface {
	SaveAll(ctx context.Context, userId uint64, cards []*pb.ProductCard) error
}

type CardService struct {
	cardDataProvider CardDataProvider
	logger           logger.Logger
	pb.UnimplementedProductCardServiceServer
}

func NewCardService(cardDataProvider CardDataProvider, logger logger.Logger) *CardService {
	return &CardService{
		cardDataProvider: cardDataProvider,
		logger:           logger,
	}
}

func (service *CardService) AddProductsCards(ctx context.Context, req *pb.AddProductsCardsRequest) (pb.Empty, error) {
	// Validate input parameters
	fmt.Println(len(req.GetProductsCards()))
	if req == nil {
		return pb.Empty{}, status.Error(codes.InvalidArgument, "request is nil")
	}
	if req.GetID() == 0 {
		return pb.Empty{}, status.Error(codes.InvalidArgument, "user ID is required")
	}
	if len(req.GetProductsCards()) == 0 {
		return pb.Empty{}, status.Error(codes.InvalidArgument, "at least one product card is required")
	}

	userId := req.GetID()
	productsCards := req.GetProductsCards()

	err := service.cardDataProvider.SaveAll(ctx, userId, productsCards)
	if err != nil {
		service.logger.Error(err)
		return pb.Empty{}, status.Errorf(codes.Internal, "could not	save cards: %v", err)
	}

	return pb.Empty{}, nil
}
