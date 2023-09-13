package response

import "backend/domain/entity"

type Payment struct {
	OrderID     string `json:"order_id"`
	UserID      string `json:"user_id"`
	PaymentType string `json:"payment_type"`
}

func NewPayment(e *entity.Payment) Payment {
	return Payment{
		UserID:      e.UserID,
		OrderID:     e.OrderID,
		PaymentType: e.PaymentType,
	}
}
