package userstorage

import (
	"context"
	"simple-service-go/modules/user/usermodel"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *usermodel.UserUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
