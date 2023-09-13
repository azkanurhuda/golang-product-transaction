package handler

import (
	"backend/domain/repository"
	"backend/domain/service"
	"backend/interfaces/middleware"
	"gorm.io/gorm"
)

type Handler struct {
	*IndexHandler
	*UserHandler
	*AuthenticationHandler
	*middleware.Middleware
	*ProductHandler
	*CartHandler
	*PaymentHandler
}

func NewHandler(db *gorm.DB, userRepository repository.User, jwtService service.Jwt, productRepository repository.Product, cartRepository repository.Cart, paymentRepository repository.Payment) *Handler {
	indexHandler := NewIndexHandler(db)
	userHandler := NewUserHandler(userRepository, jwtService)
	authenticationHandler := NewAuthenticationHandler(userRepository, jwtService)
	productHandler := NewProductHandler(productRepository, jwtService)
	cartHandler := NewCartHandler(cartRepository, jwtService)
	paymentHandler := NewPaymentHandler(paymentRepository, jwtService)

	return &Handler{
		IndexHandler:          indexHandler,
		UserHandler:           userHandler,
		AuthenticationHandler: authenticationHandler,
		Middleware:            middleware.NewHandler(jwtService),
		ProductHandler:        productHandler,
		CartHandler:           cartHandler,
		PaymentHandler:        paymentHandler,
	}
}
