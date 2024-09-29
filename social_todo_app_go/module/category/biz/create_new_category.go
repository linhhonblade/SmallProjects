package biz

import (
	"golang.org/x/net/context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/category/model"
)

// Handler -> Biz [-> Repository] -> Storage

type CreateCategoryStorage interface {
	CreateCategory(ctx context.Context, data *model.CategoryCreation) error
}

type createCategoryBiz struct {
	store CreateCategoryStorage
}

func NewCreateCategoryBiz(store CreateCategoryStorage) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) CreateNewCategory(ctx context.Context, data *model.CategoryCreation) error {
	if err := data.Validate(); err != nil {
		return err
	}
	if err := biz.store.CreateCategory(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}
	return nil
}
