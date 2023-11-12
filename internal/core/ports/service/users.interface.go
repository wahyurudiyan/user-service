package service

import (
	"context"

	"github.com/wahyurudiyan/super-sharing/user-svc/internal/core/dto/request"
	"github.com/wahyurudiyan/super-sharing/user-svc/internal/core/dto/response"
)

type UserServices interface {
	SignUp(ctx context.Context, req request.SignUp) error
	SignIn(ctx context.Context, req request.SignIn) (response.SignIn, error)

	FindUserByUniqueID(ctx context.Context, uniqueId string) (response.User, error)
	UpdateUser(ctx context.Context, req request.User) error
	DeleteUser(ctx context.Context, req request.User) error
}
