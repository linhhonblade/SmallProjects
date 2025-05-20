package query

import (
	"context"
	"github.com/google/uuid"
	sctx "github.com/linhhonblade/service-context"
	"social_todo_app_go/common"
)

type categoryByIdQuery struct {
	sctx sctx.ServiceContext
}

func NewCategoryByIdQuery(sctx sctx.ServiceContext) *categoryByIdQuery {
	return &categoryByIdQuery{sctx: sctx}
}

func (q *categoryByIdQuery) Execute(ctx context.Context, categIds []uuid.UUID) ([]CategoryDTO, error) {
	var categories []CategoryDTO

	dbCtx := q.sctx.MustGet(common.KeyGormDB).(common.DBContext)
	db := dbCtx.GetDB().Table(CategoryDTO{}.TableName())

	if len(categIds) > 0 {
		db.Where("id IN (?)", categIds)
	}

	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil

}
