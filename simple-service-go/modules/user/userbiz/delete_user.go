package userbiz

import (
	"context"
	"errors"
	"simple-service-go/modules/user/usermodel"
)

type DeleteUserStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
	SoftDeleteData(
		ctx context.Context,
		id int,
	) error
}

type deleteUserBiz struct {
	store DeleteUserStore
}

func NewDeleteUserBiz(store DeleteUserStore) *deleteUserBiz {
	return &deleteUserBiz{store: store}
}

func (biz *deleteUserBiz) DeleteUser(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if oldData.Status == 0 {
		return errors.New("data deleted")
	}
	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return err
	}
	return nil
}
