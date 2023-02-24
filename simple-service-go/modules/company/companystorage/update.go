package companystorage

import (
	"context"
	"simple-service-go/common"
	"simple-service-go/modules/company/companymodel"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *companymodel.CompanyUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
