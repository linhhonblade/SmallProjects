package usecase

import (
	"github.com/google/uuid"
	"social_todo_app_go/common"
)

type EmailPasswordRegistrationDTO struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type EmailPasswordLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponseDTO struct {
	AccessToken       string `json:"access_token"`
	AccessTokenExpIn  int    `json:"access_token_exp_in"`
	RefreshToken      string `json:"refresh_token"`
	RefreshTokenExpIn int    `json:"refresh_token_exp_in"`
}

type TokenExpResponseDTO struct {
	AccessTokenExpIn  int `json:"access_token_exp_in"`
	RefreshTokenExpIn int `json:"refresh_token_exp_in"`
}

type SetSingleImageDTO struct {
	ImageId   uuid.UUID        `json:"image_id"`
	Requester common.Requester `json:"-"`
}
