package companybiz

import (
	"context"
	"simple-service-go/common"
	"simple-service-go/modules/company/companymodel"
)

type FindCompanyStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*companymodel.Company, error)
}

type findCompanyBiz struct {
	store FindCompanyStore
}

func NewFindCompanyBiz(store FindCompanyStore) *findCompanyBiz {
	return &findCompanyBiz{store: store}
}

func (biz *findCompanyBiz) GetCompany(ctx context.Context, id int) (*companymodel.Company, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(companymodel.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(companymodel.EntityName, err)
	}
	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(companymodel.EntityName, nil)
	}
	return data, nil
}
