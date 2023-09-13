package main

import (
	"backend/infrastructure/cart"
	"backend/infrastructure/jwt"
	"backend/infrastructure/payment"
	"backend/infrastructure/product"
	"backend/infrastructure/user"
	"backend/interfaces/handler"
	"backend/interfaces/router"
	"backend/pkg/conf"
	"backend/pkg/postgres"
	"context"
	"flag"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	conn, err := postgres.NewCon()
	if err != nil {
		panic(err)
	}

	userRepository := user.NewUserRepository(conn)
	productRepository := product.NewProductRepository(conn)
	cartRepository := cart.NewCartRepository(conn)
	paymentRepository := payment.NewPaymentRepository(conn)
	jwtService := jwt.NewService()
	h := handler.NewHandler(conn, userRepository, jwtService, productRepository, cartRepository, paymentRepository)
	r := router.NewRouter(h)

	srv := &http.Server{
		Addr:         conf.Server.Addr(),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		log.Infof(" ⇨ http server started on %s", conf.Server.Addr())
		log.Infof(" ⇨ graceful timeout: %s", wait)
		if err = srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Info("received stop signal")

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer func() {
		log.Info("cancel")
		cancel()
	}()

	_ = srv.Shutdown(ctx)
	log.Infof(" ⇨ shutting down")
	os.Exit(0)
}
