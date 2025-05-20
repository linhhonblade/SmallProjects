package query

import (
	"context"
	"github.com/google/uuid"
	sctx "github.com/linhhonblade/service-context"
	"social_todo_app_go/common"
)

type getOrderByIdQuery struct {
	sctx      sctx.ServiceContext
	requester common.Requester
}

func NewGetOrderByIdQuery(sctx sctx.ServiceContext, requester common.Requester) *getOrderByIdQuery {
	return &getOrderByIdQuery{sctx: sctx, requester: requester}
}

func (q *getOrderByIdQuery) Execute(ctx context.Context, orderId uuid.UUID) (*OrderDTO, error) {
	var order *OrderDTO

	dbCtx := q.sctx.MustGet(common.KeyGormDB).(common.DBContext)
	db := dbCtx.GetDB().Table(OrderDTO{}.TableName()).Where("user_id = ?", q.requester.UserId())

	if orderId != uuid.Nil {
		db.Where("id = ?", orderId)
	}

	if err := db.Find(&order).Error; err != nil {
		return nil, err
	}

	return order, nil
}
