package companystorage

import (
	"context"
	"gorm.io/gorm"
	"simple-service-go/common"
	"simple-service-go/modules/company/companymodel"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKey ...string) (*companymodel.Company, error) {
	var data companymodel.Company
	if err := s.db.Where(conditions).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
