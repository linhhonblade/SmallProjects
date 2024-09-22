package biz

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/item/model"
)

type ListItemStorage interface {
	ListItem(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.TodoItem, error)
}

type listItemBiz struct {
	store ListItemStorage
}

func NewListItemBiz(store ListItemStorage) *listItemBiz {
	return &listItemBiz{store: store}
}

func (biz *listItemBiz) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.TodoItem, error) {
	res, err := biz.store.ListItem(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}
	return res, nil
}
