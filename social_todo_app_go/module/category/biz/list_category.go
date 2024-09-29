package biz

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/category/model"
)

type ListCategoryStorage interface {
	ListCategory(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.Category, error)
}

type listCategoryBiz struct {
	store ListCategoryStorage
}

func NewListCategoryBiz(store ListCategoryStorage) *listCategoryBiz {
	return &listCategoryBiz{store: store}
}

func (biz *listCategoryBiz) ListCategory(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Category, error) {
	res, err := biz.store.ListCategory(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}
	return res, nil
}
