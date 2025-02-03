package repository

import (
	"github.com/google/uuid"
	"social_todo_app_go/module/user/domain"
)

type UserDto struct {
	Id        uuid.UUID `gorm:"column:id;"`
	FirstName string    `gorm:"column:first_name;"`
	LastName  string    `gorm:"column:last_name;"`
	Email     string    `gorm:"column:email;"`
	Password  string    `gorm:"column:password;"`
	Salt      string    `gorm:"column:salt;"`
	Role      string    `gorm:"column:role;"`
	Status    string    `gorm:"column:status;"`
}

func (dto *UserDto) ToEntity() (*domain.User, error) {
	return domain.NewUser(
		dto.Id,
		dto.FirstName,
		dto.LastName,
		dto.Email,
		dto.Password,
		dto.Salt,
		domain.GetRole(dto.Role),
		dto.Status,
	)
}
