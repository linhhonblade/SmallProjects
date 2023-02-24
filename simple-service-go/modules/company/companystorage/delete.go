package companystorage

import (
	"context"
	"simple-service-go/common"
	"simple-service-go/modules/company/companymodel"
)

func (s *sqlStore) SoftDeleteData(
	ctx context.Context,
	id int,
) error {
	db := s.db

	if err := db.Table(companymodel.Company{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
