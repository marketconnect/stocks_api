package card_service

import (

	pb "stocks_api/app/gen/proto"
)

type CardDataProvider interface {
	GetCards(username string) ([]pb.ProductCard, error)
}