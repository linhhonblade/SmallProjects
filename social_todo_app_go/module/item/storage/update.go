package storage

import (
	"context"
	"social_todo_app_go/module/item/model"
)

func (s *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, data *model.TodoItemUpdate) error {
	if err := s.db.Where(cond).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
