package responses

import "time"

type (
	UserRegisterResponse struct {
		ID       uint64 `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		UserName string `json:"user_name"`
	}

	UserLoginResponse struct {
		Token string `json:"token"`
	}

	UserDetailResponse struct {
		ID          uint64    `json:"id"`
		Name        string    `json:"name"`
		Email       string    `json:"email"`
		UserName    string    `json:"user_name"`
		DateOfBirth time.Time `json:"date_of_birth"`
	}
)
