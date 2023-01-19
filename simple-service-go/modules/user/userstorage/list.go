package userstorage

import (
	"context"
	"simple-service-go/common"
	"simple-service-go/modules/user/usermodel"
)

func (s *sqlStore) ListDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *usermodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]usermodel.User, error) {
	var result []usermodel.User
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(usermodel.User{}.TableName()).Where(conditions).Where("status in (1)")
	if v := filter; v != nil {
		if v.Lang != "" {
			db = db.Where("lang = ?", v.Lang)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
