package card_service

import (
	"context"
	pb "stocks_api/app/gen/proto"
	"stocks_api/app/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type CardDataProvider interface {
	SaveAll(ctx context.Context, userId uint64, cards []*pb.ProductCard) (int32, error)
	GetAll(ctx context.Context, userId uint64) ([]*pb.ProductCard, error)
	Delete(ctx context.Context, userId uint64, sku uint64) error
}

type TokenManager interface {
	Verify(accessToken string) (*uint64, error)
}
type CardService struct {
	cardDataProvider CardDataProvider
	logger           logger.Logger
	tokenManager     TokenManager
	pb.UnimplementedProductCardServiceServer
}

func NewCardService(cardDataProvider CardDataProvider, tokenManager TokenManager, logger logger.Logger) *CardService {
	return &CardService{
		cardDataProvider: cardDataProvider,
		tokenManager:     tokenManager,
		logger:           logger,
	}
}

func (service *CardService) DeleteProductCard(ctx context.Context, req *pb.DeleteProductCardRequest) (*pb.DeleteProductCardResponse, error) {
	// Validate input parameters
	if req == nil {
		return &pb.DeleteProductCardResponse{}, status.Error(codes.InvalidArgument, "request is nil")
	}
	sku := req.GetSku()
	if sku == 0 {
		return &pb.DeleteProductCardResponse{}, status.Error(codes.InvalidArgument, "sku is required")
	}
	// Id extraction
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pb.DeleteProductCardResponse{}, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return &pb.DeleteProductCardResponse{}, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	userId, err := service.tokenManager.Verify(accessToken)
	if err != nil {

		return &pb.DeleteProductCardResponse{}, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}
	err = service.cardDataProvider.Delete(ctx, *userId, sku)
	if err != nil {
		service.logger.Error(err)
		return &pb.DeleteProductCardResponse{}, status.Errorf(codes.Internal, "could not delete card: %v", err)
	}
	return &pb.DeleteProductCardResponse{}, nil
}

func (service *CardService) AddProductsCards(ctx context.Context, req *pb.AddProductsCardsRequest) (*pb.AddProductsCardsResponse, error) {
	// Validate input parameters
	if req == nil {
		return &pb.AddProductsCardsResponse{}, status.Error(codes.InvalidArgument, "request is nil")
	}

	if len(req.GetProductsCards()) == 0 {
		return &pb.AddProductsCardsResponse{}, status.Error(codes.InvalidArgument, "at least one product card is required")
	}

	// Id extraction
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pb.AddProductsCardsResponse{}, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return &pb.AddProductsCardsResponse{}, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	userId, err := service.tokenManager.Verify(accessToken)
	if err != nil {

		return &pb.AddProductsCardsResponse{}, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	productsCards := req.GetProductsCards()
	// Saving
	qty, err := service.cardDataProvider.SaveAll(ctx, *userId, productsCards)
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
	// Id extraction
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pb.GetProductsCardsResponse{}, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return &pb.GetProductsCardsResponse{}, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	userId, err := service.tokenManager.Verify(accessToken)
	if err != nil {

		return &pb.GetProductsCardsResponse{}, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	productCards, err := service.cardDataProvider.GetAll(ctx, *userId)
	if err != nil {
		service.logger.Error(err)
		return &pb.GetProductsCardsResponse{}, status.Errorf(codes.Internal, "could not fetch cards: %v", err)
	}
	return &pb.GetProductsCardsResponse{ProductsCards: productCards}, nil
}
