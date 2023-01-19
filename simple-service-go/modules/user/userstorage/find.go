package userstorage

import (
	"context"
	"gorm.io/gorm"
	"simple-service-go/common"
	"simple-service-go/modules/user/usermodel"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*usermodel.User, error) {
	var data usermodel.User

	if err := s.db.Where(conditions).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
