package storage

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/item/model"
)

func (s *sqlStore) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.TodoItem, error) {
	var res []model.TodoItem
	db := s.db.Where("status <> ?", "deleted")
	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}
	if err := s.db.Table(model.TodoItem{}.TableName()).
		Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
