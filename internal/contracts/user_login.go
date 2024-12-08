package contracts

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}