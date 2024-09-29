package storage

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/category/model"
)

func (s *sqlStore) DeleteCategory(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table(model.Category{}.TableName()).Where(cond).Update("status", "deleted").Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
