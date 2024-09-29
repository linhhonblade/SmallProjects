package biz

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/category/model"
)

type UpdateCategoryStorage interface {
	GetCategory(ctx context.Context, cond map[string]interface{}) (*model.Category, error)
	UpdateCategory(ctx context.Context, cond map[string]interface{}, dataUpdate *model.CategoryUpdate) error
}

type updateCategoryBiz struct {
	store UpdateCategoryStorage
}

func NewUpdateCategoryBiz(store UpdateCategoryStorage) *updateCategoryBiz {
	return &updateCategoryBiz{store: store}
}

func (biz *updateCategoryBiz) UpdateCategoryById(ctx context.Context, id int, dataUpdate *model.CategoryUpdate) error {
	data, err := biz.store.GetCategory(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if data.Status == "deleted" {
		return common.ErrCannotUpdateEntity(model.EntityName, model.ErrCategoryDeleted)
	}
	if err := biz.store.UpdateCategory(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}
	return nil
}
