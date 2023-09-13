package handler

import (
	"backend/application/usecase"
	"backend/domain/repository"
	"backend/domain/service"
	"backend/interfaces/presenter"
	"backend/interfaces/response"
	"github.com/gorilla/mux"
	"net/http"
)

type ProductHandler struct {
	productUseCase usecase.ProductUseCase
}

func NewProductHandler(productRepository repository.Product, jwtService service.Jwt) *ProductHandler {
	return &ProductHandler{
		productUseCase: usecase.NewProductUseCase(productRepository, jwtService),
	}
}

func (h *ProductHandler) ListProductByProductCategoryID(w http.ResponseWriter, r *http.Request) {
	productCategoryID := mux.Vars(r)["product_category_id"]
	product, err := h.productUseCase.ListProductByProductCategoryID(r.Context(), productCategoryID)
	if err != nil {
		presenter.NewError(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, response.ListProducts(product))
}
