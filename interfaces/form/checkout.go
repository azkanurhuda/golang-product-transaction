package form

import (
	"backend/domain/entity"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Checkout struct {
	UserID string `form:"user_id" validate:"required"`
	CartID string `form:"cart_id" validate:"required"`
}

func (f *Checkout) Validate() error {
	validate := validator.New()
	if err := validate.Struct(f); err != nil {
		return err
	}
	return nil
}

func (f *Checkout) Entity() (*entity.Order, error) {
	return &entity.Order{
		ID:     uuid.NewString(),
		UserID: f.UserID,
		CartID: f.CartID,
	}, nil
}
