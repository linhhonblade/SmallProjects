package companybiz

import (
	"context"
	"simple-service-go/modules/company/companymodel"
)

type CreateCompanyStore interface {
	Create(ctx context.Context, data *companymodel.CompanyCreate) error
}

type createCompanyBiz struct {
	store CreateCompanyStore
}

func NewCreateCompanyBiz(store CreateCompanyStore) *createCompanyBiz {
	return &createCompanyBiz{store: store}
}

func (biz *createCompanyBiz) CreateCompany(ctx context.Context, data *companymodel.CompanyCreate) error {
	err := biz.store.Create(ctx, data)
	return err
}
