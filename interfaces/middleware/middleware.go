package middleware

import "backend/domain/service"

type Middleware struct {
	jwtService service.Jwt
}

func NewHandler(jwtService service.Jwt) *Middleware {
	return &Middleware{
		jwtService: jwtService,
	}
}
