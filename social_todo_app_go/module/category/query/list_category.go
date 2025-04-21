package query

import (
	"context"
	sctx "github.com/linhhonblade/service-context"
	"social_todo_app_go/common"
)

type listCategoryQuery struct {
	sctx sctx.ServiceContext
}

func NewListCategoryQuery(sctx sctx.ServiceContext) *listCategoryQuery {
	return &listCategoryQuery{sctx: sctx}
}

type ListCategoryFilter struct {
	Name string `form:"name" json:"name"`
}

type ListCategoryParam struct {
	common.Paging
	ListCategoryFilter
}

func (q listCategoryQuery) Execute(ctx context.Context, param *ListCategoryParam) ([]CategoryDTO, error) {
	var categories []CategoryDTO
	dbCtx := q.sctx.MustGet(common.KeyGormDB).(common.DBContext)
	db := dbCtx.GetDB().Table(CategoryDTO{}.TableName())
	if param.Name != "" {
		db.Where("name LIKE ?", "%"+param.Name+"%")
	}
	db.Count(&param.Total)
	param.Process()
	offset := param.Limit * (param.Page - 1)
	if err := db.Offset(offset).Limit(param.Limit).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
