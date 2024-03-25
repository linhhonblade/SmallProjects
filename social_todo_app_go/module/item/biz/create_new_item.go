package biz

import (
	"golang.org/x/net/context"
	"social_todo_app_go/module/item/model"
)

// Handler -> Biz [-> Repository] -> Storage

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}

type createItemBiz struct {
	store CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *createItemBiz {
	return &createItemBiz{store: store}
}

func (biz *createItemBiz) CreateNewItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := data.Validate(); err != nil {
		return err
	}
	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil
}
