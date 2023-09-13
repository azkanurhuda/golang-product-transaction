package handler

import (
	"backend/application/usecase"
	"backend/domain/repository"
	"backend/domain/service"
	"backend/interfaces/form"
	"backend/interfaces/middleware"
	"backend/interfaces/presenter"
	"backend/interfaces/response"
	"net/http"
)

type PaymentHandler struct {
	paymentUseCase usecase.PaymentUseCase
}

func NewPaymentHandler(paymentRepository repository.Payment, jwtService service.Jwt) *PaymentHandler {
	return &PaymentHandler{
		paymentUseCase: usecase.NewPaymentUseCase(paymentRepository, jwtService),
	}
}

func (h *PaymentHandler) PaymentTransaction(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		presenter.NewUnauthorized(w)
		return
	}

	f := form.Payment{
		OrderID:     r.FormValue("order_id"),
		UserID:      userID,
		PaymentType: r.FormValue("payment_type"),
	}
	if err := f.Validate(); err != nil {
		presenter.NewBadRequest(w)
		return
	}

	payment, err := f.Entity()
	if err != nil {
		presenter.NewBadRequest(w)
		return
	}

	res, err := h.paymentUseCase.PaymentTransaction(r.Context(), payment)
	if err != nil {
		presenter.NewError(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, response.NewPayment(res))
}
