package storage

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/product/model"
)

func (s *sqlStore) DeleteProduct(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table(model.Product{}.TableName()).Where(cond).Update("status", "deleted").Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
