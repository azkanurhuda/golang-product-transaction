package payment

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"context"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

var _ repository.Payment = (*PaymentRepository)(nil)

func (r *PaymentRepository) PaymentTransaction(ctx context.Context, payload *entity.Payment) error {
	if err := r.db.WithContext(ctx).Create(payload).Error; err != nil {
		return err
	}
	return nil
}
