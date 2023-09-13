package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"backend/domain/service"
	"context"
	log "github.com/sirupsen/logrus"
)

type PaymentUseCase interface {
	PaymentTransaction(ctx context.Context, payload *entity.Payment) (*entity.Payment, error)
}

type paymentUseCase struct {
	paymentRepository repository.Payment
	jwtService        service.Jwt
}

func NewPaymentUseCase(paymentRepository repository.Payment, jwtService service.Jwt) PaymentUseCase {
	return &paymentUseCase{
		paymentRepository: paymentRepository,
		jwtService:        jwtService,
	}
}

func (u *paymentUseCase) PaymentTransaction(ctx context.Context, payload *entity.Payment) (*entity.Payment, error) {
	if err := u.paymentRepository.PaymentTransaction(ctx, payload); err != nil {
		log.Error(err)
		return nil, newUnexpectedError()
	}

	res := &entity.Payment{
		ID:          payload.ID,
		OrderID:     payload.OrderID,
		UserID:      payload.UserID,
		PaymentType: payload.PaymentType,
	}

	return res, nil
}
