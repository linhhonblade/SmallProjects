package biz

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/category/model"
)

type GetCategoryStorage interface {
	GetCategory(ctx context.Context, cond map[string]interface{}) (*model.Category, error)
}

type getCategoryBiz struct {
	store GetCategoryStorage
}

func NewGetCategoryBiz(store GetCategoryStorage) *getCategoryBiz {
	return &getCategoryBiz{store: store}
}

func (biz *getCategoryBiz) GetCategoryById(ctx context.Context, id int) (*model.Category, error) {
	data, err := biz.store.GetCategory(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}
	return data, nil
}
