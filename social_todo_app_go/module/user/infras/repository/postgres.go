package repository

import (
	"context"
	"gorm.io/gorm"
	"social_todo_app_go/common"
	"social_todo_app_go/module/user/domain"
)

const TbName = "user"

type userPostgresRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) userPostgresRepo {
	return userPostgresRepo{db}
}

func (repo userPostgresRepo) Create(ctx context.Context, data *domain.User) error {
	dto := UserDto{
		Id:        data.Id(),
		FirstName: data.FirstName(),
		LastName:  data.LastName(),
		Email:     data.Email(),
		Password:  data.Password(),
		Salt:      data.Salt(),
		Role:      data.Role().String(),
	}
	if err := repo.db.Table(TbName).Create(&dto).Error; err != nil {
		return err
	}
	return nil
}

func (repo userPostgresRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var dto UserDto
	if err := repo.db.Table(TbName).Where("email = ?", email).First(&dto).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}
	return dto.ToEntity()
}
