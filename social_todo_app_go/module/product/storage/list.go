package storage

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/product/model"
)

func (s *sqlStore) ListProduct(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]model.Product, error) {
	var res []model.Product
	db := s.db.Where("status <> ?", "deleted")
	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
		if v := f.CategId; v != 0 {
			db = db.Where("categ_id = ?", v)
		}
		if v := f.Type; v != "" {
			db = db.Where("type = ?", v)
		}
	}
	if err := db.Table(model.Product{}.TableName()).
		Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&res).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return res, nil
}
