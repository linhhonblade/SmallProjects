package storage

import (
	"context"
	"gorm.io/gorm"
	"social_todo_app_go/common"
	"social_todo_app_go/module/category/model"
)

func (s *sqlStore) GetCategory(ctx context.Context, cond map[string]interface{}) (*model.Category, error) {
	var data model.Category
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
