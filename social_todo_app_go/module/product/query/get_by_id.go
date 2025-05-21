package query

import (
	"context"
	"github.com/google/uuid"
	sctx "github.com/linhhonblade/service-context"
	"social_todo_app_go/common"
)

type getProductByIdQuery struct {
	sctx sctx.ServiceContext
}

func NewGetProductByIdQuery(sctx sctx.ServiceContext) *getProductByIdQuery {
	return &getProductByIdQuery{sctx: sctx}
}

func (q *getProductByIdQuery) Execute(ctx context.Context, productIds []uuid.UUID) ([]ProductDTO, error) {
	var products []ProductDTO
	dbCtx := q.sctx.MustGet(common.KeyGormDB).(common.DBContext)
	db := dbCtx.GetDB().Table(ProductDTO{}.TableName())
	if len(productIds) > 0 {
		db.Where("id IN (?)", productIds)
	}
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
