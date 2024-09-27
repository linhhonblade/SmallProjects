package biz

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/product/model"
)

type GetProductStorage interface {
	GetProduct(ctx context.Context, cond map[string]interface{}) (*model.Product, error)
}

type getProductBiz struct {
	store GetProductStorage
}

func NewGetProductBiz(store GetProductStorage) *getProductBiz {
	return &getProductBiz{store: store}
}

func (biz *getProductBiz) GetProductById(ctx context.Context, id int) (*model.Product, error) {
	data, err := biz.store.GetProduct(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}
	return data, nil
}
