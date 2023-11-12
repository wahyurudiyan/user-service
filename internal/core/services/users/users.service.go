package users

import (
	"context"
	"time"

	"github.com/rs/xid"
	"github.com/wahyurudiyan/super-sharing/user-svc/internal/core/dto/request"
	"github.com/wahyurudiyan/super-sharing/user-svc/internal/core/dto/response"
	"github.com/wahyurudiyan/super-sharing/user-svc/internal/core/entity"
	repoIface "github.com/wahyurudiyan/super-sharing/user-svc/internal/core/ports/repository"
	serviceIface "github.com/wahyurudiyan/super-sharing/user-svc/internal/core/ports/service"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repoIface.UserRepository
}

func NewUserService(repo repoIface.UserRepository) serviceIface.UserServices {
	return &userService{
		repo: repo,
	}
}

func (s *userService) SignUp(ctx context.Context, req request.SignUp) error {
	currentTime := time.Now().UTC()
	uniqueId := xid.NewWithTime(currentTime).String()
	hashPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return err
	}

	req.Password = string(hashPass)

	userEntity := req.MapToUserEntity()
	userEntity.UniqueId = uniqueId
	userEntity.CreatedAt = currentTime
	userEntity.UpdatedAt = currentTime
	err = s.repo.CreateUser(ctx, userEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) SignIn(ctx context.Context, req request.SignIn) (response.SignIn, error) {
	user, err := s.repo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return response.SignIn{}, err
	}

	if user == (entity.User{}) {
		return response.SignIn{}, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Email), []byte(req.Password))
	if err != nil {
		return response.SignIn{}, err
	}

	// TODO: call auth service to generate token
	return response.SignIn{
		Token:    "unimplemented",
		ExpireAt: time.Now().UTC(),
	}, nil
}

func (s *userService) FindUserByUniqueID(ctx context.Context, uniqueId string) (response.User, error) {
	users, err := s.repo.FindUserByUniqueID(ctx, uniqueId)
	if err != nil {
		return response.User{}, err
	}

	respUsers := response.MapUserEntityToResponse(users)
	if len(respUsers) < 1 {
		return response.User{}, nil
	}

	return respUsers[0], nil
}

func (s *userService) UpdateUser(ctx context.Context, req request.User) error {
	user := req.MapUserRequestToUserEntity()
	err := s.repo.UpdateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteUser(ctx context.Context, req request.User) error {
	user := req.MapUserRequestToUserEntity()
	user.DeletedAt = time.Now().UTC()
	err := s.repo.DeleteUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
