package query

import (
	"context"
	"github.com/google/uuid"
	sctx "github.com/linhhonblade/service-context"
	"github.com/linhhonblade/service-context/core"
	"social_todo_app_go/common"
	productdomain "social_todo_app_go/module/product/domain"
)

type createOneProductQuery struct {
	sctx sctx.ServiceContext
}

func NewCreateOneProductQuery(sctx sctx.ServiceContext) *createOneProductQuery {
	return &createOneProductQuery{sctx: sctx}
}

func (q createOneProductQuery) Execute(ctx context.Context, data *ProductDTO) (*ProductDTO, error) {

	dbCtx := q.sctx.MustGet(common.KeyGormDB).(common.DBContext)
	db := dbCtx.GetDB().Table(productdomain.TbName)

	// business here
	if data.CategoryId == uuid.Nil {
		return nil, core.ErrBadRequest.WithError("category_id is required")
	}
	// gen primary id
	data.SQLModel = common.GenNewModel()
	// business end

	if err := db.Create(&data).Error; err != nil {
		return nil, core.ErrInternalServerError.WithError("cannot create products").WithDebug(err.Error())
	}
	return data, nil
}
