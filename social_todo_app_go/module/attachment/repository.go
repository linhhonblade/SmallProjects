package attachment

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"social_todo_app_go/common"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{db: db}
}

func (r repo) Create(ctx context.Context, att *Attachment) error {
	if err := r.db.Table(TbName).Create(att).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r repo) Find(ctx context.Context, id uuid.UUID) (*common.Attachment, error) {
	attachment := common.Attachment{}
	if err := r.db.Table(TbName).Where("id = ?", id.String()).First(&attachment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, errors.WithStack(err)
	}
	return &attachment, nil
}

func (r repo) SetAttachmentStatusActive(ctx context.Context, id uuid.UUID) error {
	_, err := r.Find(ctx, id)
	if err != nil {
		return common.ErrRecordNotFound
	}
	if err := r.db.Table(TbName).Where("id = ?", id.String()).Update("status", "active").Error; err != nil {
		return err
	}
	return nil
}
