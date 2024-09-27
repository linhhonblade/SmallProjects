package biz

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/product/model"
)

type UpdateProductStorage interface {
	GetProduct(ctx context.Context, cond map[string]interface{}) (*model.Product, error)
	UpdateProduct(ctx context.Context, cond map[string]interface{}, dataUpdate *model.ProductUpdate) error
}

type updateProductBiz struct {
	store UpdateProductStorage
}

func NewUpdateProductBiz(store UpdateProductStorage) *updateProductBiz {
	return &updateProductBiz{store: store}
}

func (biz *updateProductBiz) UpdateProductById(ctx context.Context, id int, dataUpdate *model.ProductUpdate) error {
	data, err := biz.store.GetProduct(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if data.Status == "deleted" {
		return common.ErrCannotUpdateEntity(model.EntityName, model.ErrProductDeleted)
	}
	if err := biz.store.UpdateProduct(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}
	return nil
}
