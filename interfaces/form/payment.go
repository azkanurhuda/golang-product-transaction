package form

import (
	"backend/domain/entity"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Payment struct {
	OrderID     string `form:"order_id" validate:"required"`
	UserID      string `form:"user_id" validate:"required"`
	PaymentType string `form:"payment_type" validate:"required"`
}

func (f *Payment) Validate() error {
	validate := validator.New()
	if err := validate.Struct(f); err != nil {
		return err
	}
	return nil
}

func (f *Payment) Entity() (*entity.Payment, error) {
	return &entity.Payment{
		ID:          uuid.NewString(),
		OrderID:     f.OrderID,
		UserID:      f.UserID,
		PaymentType: f.PaymentType,
	}, nil
}
