package request

import "github.com/wahyurudiyan/super-sharing/user-svc/internal/core/entity"

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname,omitempty"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
}

func (u *User) MapUserRequestToUserEntity() entity.User {
	return entity.User{
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Email:     u.Email,
		Password:  u.Password,
	}
}
