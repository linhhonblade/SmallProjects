package model

import (
	"go-hello/common"
	"go-hello/storage"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique" json:"username"`
	Password string `gorm:"not null;" json:"-"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

func (User) TableName() string {
	return "users"
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

func (user *User) Create() (*User, error) {
	err := storage.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeCreate(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := storage.Database.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(id uint) (User, error) {
	var user User
	err := storage.Database.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func ListDataByCondition(
	conditions map[string]interface{},
	filter *UserNameFilter,
	paging *common.Paging) ([]User, error) {
	var result []User
	db := storage.Database
	db = db.Table(User{}.TableName()).Where(conditions)
	if v := filter; v != nil {
		if v.Name != "" {
			db = db.Where("username ilike ? or fullname ilike ?", v.Name, v.Name)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
