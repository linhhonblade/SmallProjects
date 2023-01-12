package userbiz

import (
	"context"
	"simple-service-go/modules/user/usermodel"
)

type CreateUserStore interface {
	Create(ctx context.Context, data *usermodel.UserCreate) error
}

type createUserBiz struct {
	store CreateUserStore
}

func NewCreateUserBiz(store CreateUserStore) *createUserBiz {
	return &createUserBiz{store: store}
}

// CreateUser
// Nhiệm vụ của thằng này là làm business logic (check điều kiện mật khẩu ít nhất 4 ký tự)
// Tất cả những gì liên quan đến db vứt cho store nó làm
func (biz *createUserBiz) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}
	err := biz.store.Create(ctx, data)
	return err
}
