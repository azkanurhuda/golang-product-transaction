package response

import (
	"backend/domain/entity"
)

type Cart struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	ProductID string `json:"product_id"`
}

func NewCart(e *entity.Cart) Cart {
	return Cart{
		ID:        e.ID,
		UserID:    e.UserID,
		ProductID: e.ProductID,
	}
}

func NewListCart(e []entity.Cart) []Cart {
	var carts []Cart

	for _, cartEntity := range e {
		cart := Cart{
			ID:        cartEntity.ID,
			UserID:    cartEntity.UserID,
			ProductID: cartEntity.ProductID,
		}
		carts = append(carts, cart)
	}

	return carts
}

type Checkout struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	ChartID string `json:"cart_id"`
}

func NewCheckout(e *entity.Order) Checkout {
	return Checkout{
		ID:      e.ID,
		UserID:  e.UserID,
		ChartID: e.CartID,
	}
}
