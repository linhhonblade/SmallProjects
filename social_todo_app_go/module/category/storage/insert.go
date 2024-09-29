package storage

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/category/model"
)

func (s *sqlStore) CreateCategory(ctx context.Context, data *model.CategoryCreation) error {
	data.SQLModel = common.GenNewModel()
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
