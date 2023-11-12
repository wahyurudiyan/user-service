package response

import (
	"time"

	"github.com/wahyurudiyan/super-sharing/user-svc/internal/core/entity"
)

type User struct {
	UniqueId  string    `json:"unique_id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname,omitempty"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Users []User

func MapUserEntityToResponse(users entity.Users) Users {
	var usersResp Users
	for _, u := range users {
		usersResp = append(usersResp, User{
			UniqueId:  u.UniqueId,
			Firstname: u.Firstname,
			Lastname:  u.Lastname,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}

	return usersResp
}
