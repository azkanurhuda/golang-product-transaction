package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"backend/domain/service"
	"context"
	log "github.com/sirupsen/logrus"
)

type CartUseCase interface {
	AddProductToCart(ctx context.Context, cart *entity.Cart) (*entity.Cart, error)
	GetCartByUserID(ctx context.Context, userID string) ([]entity.Cart, error)
	DeleteCartByID(ctx context.Context, id string) error
	CheckoutCartByUserID(ctx context.Context, payload *entity.Order) (*entity.Order, error)
}

type cartUseCase struct {
	cartRepository repository.Cart
	jwtService     service.Jwt
}

func NewCartUseCase(cartRepository repository.Cart, jwtService service.Jwt) CartUseCase {
	return &cartUseCase{
		cartRepository: cartRepository,
		jwtService:     jwtService,
	}
}

func (u *cartUseCase) AddProductToCart(ctx context.Context, cart *entity.Cart) (*entity.Cart, error) {
	if err := u.cartRepository.AddProductToCart(ctx, cart); err != nil {
		log.Error(err)
		return nil, newUnexpectedError()
	}

	res := &entity.Cart{
		ID:        cart.ID,
		UserID:    cart.UserID,
		ProductID: cart.ProductID,
	}

	return res, nil
}

func (u *cartUseCase) GetCartByUserID(ctx context.Context, userID string) ([]entity.Cart, error) {
	cart, err := u.cartRepository.GetCartByUserID(ctx, userID)
	if err != nil {
		log.Error(err)
		return nil, newUnexpectedError()
	}

	return cart, nil
}

func (u *cartUseCase) DeleteCartByID(ctx context.Context, id string) error {
	err := u.cartRepository.DeleteCartByID(ctx, id)
	if err != nil {
		log.Error(err)
		return newUnexpectedError()
	}

	return nil
}

func (u *cartUseCase) CheckoutCartByUserID(ctx context.Context, payload *entity.Order) (*entity.Order, error) {
	if err := u.cartRepository.CheckoutCartByUserID(ctx, payload); err != nil {
		log.Error(err)
		return nil, newUnexpectedError()
	}

	res := &entity.Order{
		ID:     payload.ID,
		CartID: payload.CartID,
		UserID: payload.UserID,
	}

	return res, nil
}
