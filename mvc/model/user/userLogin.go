package user

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}