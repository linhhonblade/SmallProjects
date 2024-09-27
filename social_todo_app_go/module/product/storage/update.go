package storage

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/product/model"
)

func (s *sqlStore) UpdateProduct(ctx context.Context, cond map[string]interface{}, data *model.ProductUpdate) error {
	if err := s.db.Model(&model.Product{}).Where(cond).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
