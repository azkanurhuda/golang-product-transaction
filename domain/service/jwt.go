package service

import "backend/domain/entity"

type Jwt interface {
	Sign(user *entity.User) (*entity.AccessToken, error)
	Verify(signedToken string) (string, error)
}
