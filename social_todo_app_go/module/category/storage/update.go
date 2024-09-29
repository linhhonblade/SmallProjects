package storage

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/category/model"
)

func (s *sqlStore) UpdateCategory(ctx context.Context, cond map[string]interface{}, data *model.CategoryUpdate) error {
	if err := s.db.Where(cond).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
