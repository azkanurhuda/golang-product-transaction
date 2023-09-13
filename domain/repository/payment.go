package repository

import (
	"backend/domain/entity"
	"context"
)

type Payment interface {
	PaymentTransaction(ctx context.Context, payload *entity.Payment) error
}
