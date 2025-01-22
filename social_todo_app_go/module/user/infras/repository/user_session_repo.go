package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"social_todo_app_go/module/user/domain"
	"time"
)

const TbUserSessionName = "user_session"

type sessionPostgresRepo struct {
	db *gorm.DB
}

func NewUserSessionPostgresRepo(db *gorm.DB) sessionPostgresRepo {
	return sessionPostgresRepo{db}
}

func (repo sessionPostgresRepo) Create(ctx context.Context, data *domain.UserSession) error {
	dto := userSessionDto{
		Id:           data.Id(),
		UserId:       data.UserId(),
		RefreshToken: data.RefreshToken(),
		AccessExpAt:  data.AccessExpAt(),
		RefreshExpAt: data.RefreshExpAt(),
	}
	return repo.db.Table(TbUserSessionName).Create(&dto).Error
}

type userSessionDto struct {
	Id           uuid.UUID `gorm:"column:id"`
	UserId       uuid.UUID `gorm:"column:user_id;"`
	RefreshToken string    `gorm:"column:refresh_token;"`
	AccessExpAt  time.Time `gorm:"column:access_expire_date;"`
	RefreshExpAt time.Time `gorm:"column:refresh_expire_date;"`
}
