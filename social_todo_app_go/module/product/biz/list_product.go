package biz

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/product/model"
)

type ListProductStorage interface {
	ListProduct(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.Product, error)
}

type listProductBiz struct {
	store ListProductStorage
}

func NewListProductBiz(store ListProductStorage) *listProductBiz {
	return &listProductBiz{store: store}
}

func (biz *listProductBiz) ListProduct(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Product, error) {
	res, err := biz.store.ListProduct(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}
	return res, nil
}
