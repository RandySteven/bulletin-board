package responses

type UserRegisterResponse struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
