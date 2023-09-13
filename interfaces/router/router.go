package router

import (
	"backend/interfaces/handler"
	"backend/interfaces/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(h *handler.Handler) *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.Logging)
	r.Use(middleware.CORS)

	r.HandleFunc("/", h.Index).Methods("GET")
	r.HandleFunc("/healthy", h.Healthy).Methods("GET")
	r.HandleFunc("/signup", h.Signup).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")

	v1 := r.PathPrefix("/v1").Subrouter()
	v1.Use(h.JWT())
	v1.HandleFunc("/me", h.Me).Methods("GET")
	v1.HandleFunc("/products/{product_category_id}", h.ListProductByProductCategoryID).Methods("GET")
	v1.HandleFunc("/carts", h.AddProductToCart).Methods("POST")
	v1.HandleFunc("/carts", h.GetCartByUserID).Methods("GET")
	v1.HandleFunc("/carts/{id}", h.DeleteCartByID).Methods("DELETE")
	v1.HandleFunc("/carts/checkout", h.CheckoutCartByUserID).Methods("POST")
	v1.HandleFunc("/payment/transaction", h.PaymentTransaction).Methods("POST")
	return r
}
