package repository

import (
	"context"

	"github.com/wahyurudiyan/super-sharing/user-svc/internal/core/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user entity.User) error
	FindUsers(ctx context.Context, limit uint) (entity.Users, error)
	FindUserByEmail(ctx context.Context, email string) (entity.User, error)
	FindUserByUniqueID(ctx context.Context, uniqueId string) (entity.Users, error)
	UpdateUser(ctx context.Context, user entity.User) error
	DeleteUser(ctx context.Context, user entity.User) error
}
