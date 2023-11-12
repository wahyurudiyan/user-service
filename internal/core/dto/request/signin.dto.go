package request

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
