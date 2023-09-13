package cart

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"context"
	"errors"
	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

var _ repository.Cart = (*CartRepository)(nil)

func (r *CartRepository) AddProductToCart(ctx context.Context, cart *entity.Cart) error {
	if err := r.db.WithContext(ctx).Create(cart).Error; err != nil {
		return err
	}
	return nil
}

func (r *CartRepository) GetCartByUserID(ctx context.Context, userID string) ([]entity.Cart, error) {
	var cart []entity.Cart

	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return cart, nil
}

func (r *CartRepository) DeleteCartByID(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Cart{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	return nil
}

func (r *CartRepository) CheckoutCartByUserID(ctx context.Context, checkoutCart *entity.Order) error {
	if err := r.db.WithContext(ctx).Create(checkoutCart).Error; err != nil {
		return err
	}

	return nil
}
