package storage

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/product/model"
)

func (s *sqlStore) CreateProduct(ctx context.Context, data *model.ProductCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
