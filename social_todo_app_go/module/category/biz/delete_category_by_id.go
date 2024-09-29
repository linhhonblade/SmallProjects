package biz

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/category/model"
)

type DeleteCategoryStorage interface {
	GetCategory(ctx context.Context, cond map[string]interface{}) (*model.Category, error)
	DeleteCategory(ctx context.Context, cond map[string]interface{}) error
}

type deleteCategoryBiz struct {
	store DeleteCategoryStorage
}

func NewDeleteCategoryBiz(store DeleteCategoryStorage) *deleteCategoryBiz {
	return &deleteCategoryBiz{store: store}
}

func (biz *deleteCategoryBiz) DeleteCategoryById(ctx context.Context, id int) error {
	data, err := biz.store.GetCategory(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(model.EntityName, err)
	}
	if data.Status == "deleted" {
		return common.ErrCannotDeleteEntity(model.EntityName, model.ErrCategoryDeleted)
	}
	if err := biz.store.DeleteCategory(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}
	return nil
}
