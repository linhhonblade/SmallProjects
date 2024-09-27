package biz

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/product/model"
)

type DeleteProductStorage interface {
	GetProduct(ctx context.Context, cond map[string]interface{}) (*model.Product, error)
	DeleteProduct(ctx context.Context, cond map[string]interface{}) error
}

type deleteProductBiz struct {
	store DeleteProductStorage
}

func NewDeleteProductBiz(store DeleteProductStorage) *deleteProductBiz {
	return &deleteProductBiz{store: store}
}

func (biz *deleteProductBiz) DeleteProductById(ctx context.Context, id int) error {
	data, err := biz.store.GetProduct(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(model.EntityName, err)
	}
	if data.Status == "deleted" {
		return common.ErrCannotDeleteEntity(model.EntityName, model.ErrProductDeleted)
	}
	if err := biz.store.DeleteProduct(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}
	return nil
}
