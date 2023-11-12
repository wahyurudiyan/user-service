package response

import "time"

type SignIn struct {
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expire_at"`
}
