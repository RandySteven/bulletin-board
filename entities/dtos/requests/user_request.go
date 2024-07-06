package requests

import "task_mission/enums"

type (
	UserRegisterRequest struct {
		FirstName   string           `json:"first_name" validate:"required"`
		LastName    string           `json:"last_name" validate:"required"`
		UserName    string           `json:"user_name" validate:"required"`
		DateOfBirth string           `json:"date_of_birth" validate:"date_of_birth"`
		Gender      enums.UserGender `json:"gender" validate:"required"`
		Email       string           `json:"email" validate:"email"`
		Password    string           `json:"password" validate:"password"`
	}

	UserLoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
