package request

import "github.com/wahyurudiyan/super-sharing/user-svc/internal/core/entity"

type SignUp struct {
	Firstname string `validate:"required"json:"firstname"`
	Lastname  string `validate:"required"json:"lastname"`
	Email     string `validate:"required,email"json:"email"`
	Password  string `validate:"required"json:"password"`
}

func (s SignUp) MapToUserEntity() entity.User {
	return entity.User{
		Firstname: s.Firstname,
		Lastname:  s.Lastname,
		Email:     s.Email,
		Password:  s.Password,
	}
}
