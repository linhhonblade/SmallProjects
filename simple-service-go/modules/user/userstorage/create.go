package userstorage

import (
	"context"
	"simple-service-go/common"
	"simple-service-go/modules/user/usermodel"
)

func (s *sqlStore) Create(ctx context.Context, data *usermodel.UserCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
