package storage

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/category/model"
)

func (s *sqlStore) ListCategory(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Category, error) {
	var res []model.Category
	db := s.db.Where("status <> ?", "deleted")
	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}
	if err := s.db.Table(model.Category{}.TableName()).
		Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	if err := db.Debug().Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&res).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return res, nil
}
