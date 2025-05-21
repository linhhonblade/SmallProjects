package query

import (
	"context"
	"github.com/google/uuid"
	sctx "github.com/linhhonblade/service-context"
	"github.com/linhhonblade/service-context/core"
	"social_todo_app_go/common"
)

type listOrderQuery struct {
	sctx        sctx.ServiceContext
	requester   common.Requester
	productRepo ProductRepository
}

func NewListOrderQuery(sctx sctx.ServiceContext, requester common.Requester, productRepo ProductRepository) *listOrderQuery {
	return &listOrderQuery{sctx: sctx, requester: requester, productRepo: productRepo}
}

type ListOrderFilter struct {
	Status string `form:"status" json:"status"`
}

type ListOrderParam struct {
	common.Paging
	ListOrderFilter
}

func (q *listOrderQuery) Execute(ctx context.Context, param *ListOrderParam) ([]OrderDTO, error) {
	var orders []OrderDTO
	dbCtx := q.sctx.MustGet(common.KeyGormDB).(common.DBContext)
	db := dbCtx.GetDB().Table(OrderDTO{}.TableName()).Where("user_id = ?", q.requester.UserId())

	if param.Status != "" {
		db.Where("status = ?", param.Status)
	}

	db.Count(&param.Total)
	param.Process()
	offset := param.Limit * (param.Page - 1)
	// Get order list together with order items, only preload order item
	db = db.Preload("OrderItems")
	if err := db.Offset(offset).Limit(param.Limit).Find(&orders).Error; err != nil {
		return nil, err
	}

	productIds := []uuid.UUID{}
	for i := range orders {
		for j := range orders[i].OrderItems {
			productIds = append(productIds, orders[i].OrderItems[j].ProductId)
		}
	}
	// Get product data via grpc
	products, err := q.productRepo.FindWithIds(ctx, productIds)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("cannot get product of order item").WithDebug(err.Error())
	}
	productMap := make(map[uuid.UUID]*ProductDTO)
	for _, p := range products {
		productMap[p.Id] = &p
	}

	// Parse product data to order item
	for i := range orders {
		for j := range orders[i].OrderItems {
			orders[i].OrderItems[j].Product = productMap[orders[i].OrderItems[j].ProductId]
		}
	}
	return orders, nil
}

type ProductRepository interface {
	FindWithIds(ctx context.Context, ids []uuid.UUID) ([]ProductDTO, error)
}
