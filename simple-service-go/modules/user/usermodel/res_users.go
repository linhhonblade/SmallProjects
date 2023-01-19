package usermodel

import (
	"errors"
	"simple-service-go/common"
	"strings"
)

type User struct {
	common.SQLModel `json:",inline"`
	Login           string `json:"login" gorm:"column:login"`
	Password        string `json:"password" gorm:"column:password"`
	Lang            string `json:"lang" gorm:"column:lang"`
}

func (User) TableName() string {
	return "res_users"
}

type UserUpdate struct {
	Login    *string `json:"login" gorm:"column:login"`
	Password *string `json:"password" gorm:"column:password"`
	Lang     *string `json:"lang" gorm:"column:lang"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

type UserCreate struct {
	Id       int    `json:"id" gorm:"column:id"`
	Login    string `json:"login" gorm:"column:login"`
	Password string `json:"password" gorm:"column:password"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (user *UserCreate) Validate() error {
	user.Login = strings.TrimSpace(user.Login)
	if len(user.Login) == 0 {
		return errors.New("user login cannot be blank")
	}
	return nil
}
