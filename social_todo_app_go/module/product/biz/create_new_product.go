package biz

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/product/model"
)

type CreateProductStorage interface {
	CreateProduct(ctx context.Context, data *model.ProductCreation) error
}

type createProductBiz struct {
	store CreateProductStorage
}

func NewCreateProductBiz(store CreateProductStorage) *createProductBiz {
	return &createProductBiz{store: store}
}

func (biz *createProductBiz) CreateNewProduct(ctx context.Context, data *model.ProductCreation) error {
	if err := biz.store.CreateProduct(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}
	return nil
}
