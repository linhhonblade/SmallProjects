package query

import (
	"context"
	sctx "github.com/linhhonblade/service-context"
	"github.com/linhhonblade/service-context/core"
	"social_todo_app_go/common"
)

type CreateOneProductCategoryQuery struct {
	sctx sctx.ServiceContext
}

func NewCreateOneProductCategoryQuery(sctx sctx.ServiceContext) *CreateOneProductCategoryQuery {
	return &CreateOneProductCategoryQuery{sctx: sctx}
}

func (q *CreateOneProductCategoryQuery) Execute(ctx context.Context, data *CategoryDTO) (*CategoryDTO, error) {
	dbCtx := q.sctx.MustGet(common.KeyGormDB).(common.DBContext)
	db := dbCtx.GetDB().Table(CategoryDTO{}.TableName())

	// business here
	if data.Name == "" {
		return nil, core.ErrBadRequest.WithError("name is required")
	}

	// gen primary id
	data.Id = common.GenUUID().String()
	// business end

	if err := db.Create(&data).Error; err != nil {
		return nil, core.ErrInternalServerError.WithError("cannot create products").WithDebug(err.Error())
	}
	return data, nil
}
