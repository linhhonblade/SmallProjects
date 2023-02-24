package companybiz

import (
	"context"
	"errors"
	"simple-service-go/modules/company/companymodel"
)

type UpdateCompanyStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKey ...string,
	) (*companymodel.Company, error)
	UpdateData(
		ctx context.Context,
		id int,
		data *companymodel.CompanyUpdate,
	) error
}

type updateCompanyBiz struct {
	store UpdateCompanyStore
}

func NewUpdateCompanyBiz(store UpdateCompanyStore) *updateCompanyBiz {
	return &updateCompanyBiz{store: store}
}

func (biz *updateCompanyBiz) UpdateData(ctx context.Context, id int, data *companymodel.CompanyUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if oldData.Status == 0 {
		return errors.New("data deleted")
	}
	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return err
	}
	return nil
}
