package query

import (
	"context"
	"github.com/google/uuid"
	sctx "github.com/linhhonblade/service-context"
	"github.com/linhhonblade/service-context/core"
	"social_todo_app_go/common"
)

type getOrderByIdQuery struct {
	sctx        sctx.ServiceContext
	requester   common.Requester
	productRepo ProductRepository
}

func NewGetOrderByIdQuery(sctx sctx.ServiceContext, requester common.Requester, productRepo ProductRepository) *getOrderByIdQuery {
	return &getOrderByIdQuery{sctx: sctx, requester: requester, productRepo: productRepo}
}

func (q *getOrderByIdQuery) Execute(ctx context.Context, orderId uuid.UUID) (*OrderDTO, error) {
	var order *OrderDTO

	dbCtx := q.sctx.MustGet(common.KeyGormDB).(common.DBContext)
	db := dbCtx.GetDB().Table(OrderDTO{}.TableName()).Where("user_id = ?", q.requester.UserId())

	if orderId != uuid.Nil {
		db.Where("id = ?", orderId)
	}

	db.Preload("OrderItems")
	if err := db.Find(&order).Error; err != nil {
		return nil, err
	}

	productIds := []uuid.UUID{}
	for i := range order.OrderItems {
		productIds = append(productIds, order.OrderItems[i].ProductId)
	}

	products, err := q.productRepo.FindWithIds(ctx, productIds)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("cannot get product of order item").WithDebug(err.Error())
	}
	productMap := make(map[uuid.UUID]*ProductDTO)
	for _, p := range products {
		productMap[p.Id] = &p
	}

	for i := range order.OrderItems {
		order.OrderItems[i].Product = productMap[order.OrderItems[i].ProductId]
	}

	return order, nil
}
