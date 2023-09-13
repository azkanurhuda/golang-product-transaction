package handler

import (
	"backend/application/usecase"
	"backend/domain/repository"
	"backend/domain/service"
	"backend/interfaces/form"
	"backend/interfaces/middleware"
	"backend/interfaces/presenter"
	"backend/interfaces/response"
	"github.com/gorilla/mux"
	"net/http"
)

type CartHandler struct {
	cartUseCase usecase.CartUseCase
}

func NewCartHandler(cartRepository repository.Cart, jwtService service.Jwt) *CartHandler {
	return &CartHandler{
		cartUseCase: usecase.NewCartUseCase(cartRepository, jwtService),
	}
}

func (h *CartHandler) AddProductToCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		presenter.NewUnauthorized(w)
		return
	}

	f := form.Cart{
		UserID:    userID,
		ProductID: r.FormValue("product_id"),
	}
	if err := f.Validate(); err != nil {
		presenter.NewBadRequest(w)
		return
	}

	cart, err := f.Entity()
	if err != nil {
		presenter.NewBadRequest(w)
		return
	}

	res, err := h.cartUseCase.AddProductToCart(r.Context(), cart)
	if err != nil {
		presenter.NewError(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, response.NewCart(res))
}

func (h *CartHandler) GetCartByUserID(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		presenter.NewUnauthorized(w)
		return
	}

	cart, err := h.cartUseCase.GetCartByUserID(r.Context(), userID)
	if err != nil {
		presenter.NewError(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, response.NewListCart(cart))
}

func (h *CartHandler) DeleteCartByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.cartUseCase.DeleteCartByID(r.Context(), id)
	if err != nil {
		presenter.NewError(w, err)
		return
	}

	presenter.OK(w)
}

func (h *CartHandler) CheckoutCartByUserID(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		presenter.NewUnauthorized(w)
		return
	}

	f := form.Checkout{
		UserID: userID,
		CartID: r.FormValue("cart_id"),
	}
	if err := f.Validate(); err != nil {
		presenter.NewBadRequest(w)
		return
	}

	checkout, err := f.Entity()
	if err != nil {
		presenter.NewBadRequest(w)
		return
	}

	res, err := h.cartUseCase.CheckoutCartByUserID(r.Context(), checkout)
	if err != nil {
		presenter.NewError(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, response.NewCheckout(res))
}
