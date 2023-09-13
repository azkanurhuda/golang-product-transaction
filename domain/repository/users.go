package repository

import (
	"backend/domain/entity"
	"context"
)

type User interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByID(ctx context.Context, userID string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, userID string) error
}
