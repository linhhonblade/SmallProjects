package query

import (
	"context"
	"github.com/google/uuid"
	sctx "github.com/linhhonblade/service-context"
	"github.com/linhhonblade/service-context/core"
	"social_todo_app_go/common"
	productdomain "social_todo_app_go/module/product/domain"
)

type ProductDTO struct {
	common.SQLModel
	Name       string       `json:"name" gorm:"column:name;"`
	Type       string       `json:"type" gorm:"column:type;"`
	CategoryId uuid.UUID    `json:"category_id" gorm:"column:category_id;"`
	Category   *CategoryDTO `json:"category" gorm:"refere"`
}

type CategoryDTO struct {
	Id   uuid.UUID `json:"id" gorm:"column:id;"`
	Name string    `json:"name" gorm:"column:name;"`
}

func (CategoryDTO) TableName() string {
	return "category"
}

type listProductQuery struct {
	sctx sctx.ServiceContext
}

func NewListProductQuery(sctx sctx.ServiceContext) *listProductQuery {
	return &listProductQuery{sctx: sctx}
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
	db = db.Preload("Category")
	offset := param.Limit * (param.Page - 1)

	if err := db.Offset(offset).Limit(param.Limit).Order("id desc").Find(&products).Error; err != nil {
		return nil, core.ErrInternalServerError.WithError("cannot list products").WithDebug(err.Error())
	}
	return products, nil
}
