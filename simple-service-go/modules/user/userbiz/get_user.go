package userbiz

import (
	"context"
	"errors"
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

func (biz *findUserBiz) FindUser(ctx context.Context, id int) (*usermodel.User, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	if data.Status == 0 {
		return nil, errors.New("record deleted")
	}
	return data, nil
}
