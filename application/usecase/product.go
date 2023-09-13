package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"backend/domain/service"
	"context"
	log "github.com/sirupsen/logrus"
)

type ProductUseCase interface {
	ListProductByProductCategoryID(ctx context.Context, productCategoryID string) ([]entity.Product, error)
}

type productUseCase struct {
	productRepository repository.Product
	jwtService        service.Jwt
}

func NewProductUseCase(productRepository repository.Product, jwtService service.Jwt) ProductUseCase {
	return &productUseCase{
		productRepository: productRepository,
		jwtService:        jwtService,
	}
}

func (u *productUseCase) ListProductByProductCategoryID(ctx context.Context, productCategoryID string) ([]entity.Product, error) {
	product, err := u.productRepository.GetProductByProductCategoryID(ctx, productCategoryID)
	if err != nil {
		log.Error(err)
		return nil, newUnexpectedError()
	}

	return product, nil
}
