package card_service

import (
	"context"
	pb "stocks_api/app/gen/proto"
	"stocks_api/app/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CardDataProvider interface {
	SaveAll(ctx context.Context, userId uint64, cards []*pb.ProductCard) (int32, error)
	GetAll(ctx context.Context, userId uint64) ([]*pb.ProductCard, error)
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

func (service *CardService) AddProductsCards(ctx context.Context, req *pb.AddProductsCardsRequest) (*pb.AddProductsCardsResponse, error) {
	// Validate input parameters
	if req == nil {
		return &pb.AddProductsCardsResponse{}, status.Error(codes.InvalidArgument, "request is nil")
	}
	if req.GetID() == 0 {
		return &pb.AddProductsCardsResponse{}, status.Error(codes.InvalidArgument, "user ID is required")
	}
	if len(req.GetProductsCards()) == 0 {
		return &pb.AddProductsCardsResponse{}, status.Error(codes.InvalidArgument, "at least one product card is required")
	}

	userId := req.GetID()
	productsCards := req.GetProductsCards()

	qty, err := service.cardDataProvider.SaveAll(ctx, userId, productsCards)
	if err != nil {
		service.logger.Error(err)
		return &pb.AddProductsCardsResponse{}, status.Errorf(codes.Internal, "could not	save cards: %v", err)
	}

	return &pb.AddProductsCardsResponse{Qty: qty}, nil
}

func (service *CardService) GetProductsCards(ctx context.Context, req *pb.GetProductsCardsRequest) (*pb.GetProductsCardsResponse, error) {
	// Validate input parameters
	if req == nil {
		return &pb.GetProductsCardsResponse{}, status.Error(codes.InvalidArgument, "request is nil")
	}
	if req.GetID() == 0 {
		return &pb.GetProductsCardsResponse{}, status.Error(codes.InvalidArgument, "user ID is required")
	}
	productCards, err := service.cardDataProvider.GetAll(ctx, req.GetID())
	if err != nil {
		service.logger.Error(err)
		return &pb.GetProductsCardsResponse{}, status.Errorf(codes.Internal, "could not fetch cards: %v", err)
	}
	return &pb.GetProductsCardsResponse{ProductsCards: productCards}, nil
}
