package userbiz

import (
	"context"
	"errors"
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
	if len(data.Password) < 4 {
		return errors.New("Password must be at least 4 characters.")
	}
	err := biz.store.Create(ctx, data)
	return err
}
