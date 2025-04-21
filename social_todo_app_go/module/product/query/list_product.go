package query

import (
	"context"
	"github.com/google/uuid"
	sctx "github.com/linhhonblade/service-context"
	"github.com/linhhonblade/service-context/core"
	"social_todo_app_go/common"
	productdomain "social_todo_app_go/module/product/domain"
)

type listProductQuery struct {
	sctx      sctx.ServiceContext
	categRepo CategoryRepository
}

func NewListProductQuery(sctx sctx.ServiceContext, categRepo CategoryRepository) *listProductQuery {
	return &listProductQuery{sctx: sctx, categRepo: categRepo}
}

type ListProductFilter struct {
	CategoryId string `form:"category_id" json:"category_id"`
}

type ListProductParam struct {
	common.Paging
	ListProductFilter
}

func (q listProductQuery) Execute(ctx context.Context, param *ListProductParam) ([]ProductDTO, error) {
	var products []ProductDTO

	dbCtx := q.sctx.MustGet(common.KeyGormDB).(common.DBContext)
	db := dbCtx.GetDB().Table(productdomain.TbName)

	if param.CategoryId != "" {
		db.Where("category_id = ?", param.CategoryId)
	}

	db.Count(&param.Total)
	param.Process()

	//db = db.Preload("Category")

	offset := param.Limit * (param.Page - 1)

	if err := db.Offset(offset).Limit(param.Limit).Order("id desc").Find(&products).Error; err != nil {
		return nil, core.ErrInternalServerError.WithError("cannot list products").WithDebug(err.Error())
	}

	catIds := []uuid.UUID{}
	categMap := make(map[uuid.UUID]*CategoryDTO)

	for i := range products {
		catIds = append(catIds, products[i].CategoryId)
	}

	categories, err := q.categRepo.FindWithIds(ctx, catIds)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("cannot list products").WithDebug(err.Error())
	}

	for i := range categories {
		categMap[categories[i].Id] = &categories[i]
	}

	for i := range products {
		products[i].Category = categMap[products[i].CategoryId]
	}

	return products, nil
}

type CategoryRepository interface {
	FindWithIds(ctx context.Context, ids []uuid.UUID) ([]CategoryDTO, error)
}
