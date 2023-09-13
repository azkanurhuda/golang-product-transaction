package product

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"context"
	"errors"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

var _ repository.Product = (*ProductRepository)(nil)

func (r *ProductRepository) GetProductByProductCategoryID(ctx context.Context, productCategoryID string) ([]entity.Product, error) {
	var product []entity.Product

	if err := r.db.WithContext(ctx).Where("product_category_id = ?", productCategoryID).Find(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return product, nil
}
