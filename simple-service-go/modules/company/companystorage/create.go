package companystorage

import (
	"context"
	"simple-service-go/common"
	"simple-service-go/modules/company/companymodel"
)

func (s *sqlStore) Create(ctx context.Context, data *companymodel.CompanyCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
