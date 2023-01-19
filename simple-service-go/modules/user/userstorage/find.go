package userstorage

import (
	"context"
	"simple-service-go/modules/user/usermodel"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*usermodel.User, error) {
	var data usermodel.User

	if err := s.db.Where(conditions).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
