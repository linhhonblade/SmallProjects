package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"social_todo_app_go/common"
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

func (repo sessionPostgresRepo) Find(ctx context.Context, sessionId uuid.UUID) (*domain.UserSession, error) {
	var dto userSessionDto
	if err := repo.db.Table(TbUserSessionName).Where("id = ?", sessionId.String()).First(&dto).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}
	return dto.ToEntity()
}

func (repo sessionPostgresRepo) Delete(ctx context.Context, sessionId uuid.UUID) error {
	if err := repo.db.Table(TbUserSessionName).Where("id = ?", sessionId.String()).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}

func (repo sessionPostgresRepo) FindByRefreshToken(ctx context.Context, refreshToken string) (*domain.UserSession, error) {
	var dto userSessionDto
	if err := repo.db.Table(TbUserSessionName).Where("refresh_token = ?", refreshToken).First(&dto).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}
	return dto.ToEntity()
}

type userSessionDto struct {
	Id           uuid.UUID `gorm:"column:id"`
	UserId       uuid.UUID `gorm:"column:user_id;"`
	RefreshToken string    `gorm:"column:refresh_token;"`
	AccessExpAt  time.Time `gorm:"column:access_expire_date;"`
	RefreshExpAt time.Time `gorm:"column:refresh_expire_date;"`
}

func (dto *userSessionDto) ToEntity() (*domain.UserSession, error) {
	return domain.NewUserSession(
		dto.Id,
		dto.UserId,
		dto.RefreshToken,
		dto.AccessExpAt,
		dto.RefreshExpAt,
	), nil
}
