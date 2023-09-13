package repository

import (
	"backend/domain/entity"
	"context"
)

type Product interface {
	GetProductByProductCategoryID(ctx context.Context, productCategoryID string) ([]entity.Product, error)
}
