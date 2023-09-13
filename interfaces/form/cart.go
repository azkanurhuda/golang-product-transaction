package form

import (
	"backend/domain/entity"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Cart struct {
	UserID    string `form:"user_id" validate:"required"`
	ProductID string `form:"product_id" validate:"required"`
}

func (f *Cart) Validate() error {
	validate := validator.New()
	if err := validate.Struct(f); err != nil {
		return err
	}
	return nil
}

func (f *Cart) Entity() (*entity.Cart, error) {
	return &entity.Cart{
		ID:        uuid.NewString(),
		UserID:    f.UserID,
		ProductID: f.ProductID,
	}, nil
}
