package repository

import (
	"backend/domain/entity"
	"context"
)

type Cart interface {
	AddProductToCart(ctx context.Context, cart *entity.Cart) error
	GetCartByUserID(ctx context.Context, userID string) ([]entity.Cart, error)
	DeleteCartByID(ctx context.Context, id string) error
	CheckoutCartByUserID(ctx context.Context, checkoutCart *entity.Order) error
}
