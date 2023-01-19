package userbiz

import (
	"context"
	"simple-service-go/common"
	"simple-service-go/modules/user/usermodel"
)

type FindUserStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
}

type findUserBiz struct {
	store FindUserStore
}

func NewFindUserBiz(store FindUserStore) *findUserBiz {
	return &findUserBiz{store: store}
}

func (biz *findUserBiz) GetUser(ctx context.Context, id int) (*usermodel.User, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}
	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(usermodel.EntityName, nil)
	}
	return data, nil
}
