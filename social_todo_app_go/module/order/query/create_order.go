package query

import (
	"context"
	sctx "github.com/linhhonblade/service-context"
	"github.com/linhhonblade/service-context/core"
	"social_todo_app_go/common"
	"social_todo_app_go/module/sequence"
)

type createOneOrderQuery struct {
	sctx      sctx.ServiceContext
	requester common.Requester
}

func NewCreateOneOrderQuery(sctx sctx.ServiceContext, requester common.Requester) *createOneOrderQuery {
	return &createOneOrderQuery{sctx: sctx, requester: requester}
}

func (q *createOneOrderQuery) Execute(ctx context.Context, order *OrderDTO) (*OrderDTO, error) {
	dbCtx := q.sctx.MustGet(common.KeyGormDB).(common.DBContext)
	db := dbCtx.GetDB().Table(OrderDTO{}.TableName())

	// business here
	order.UserId = q.requester.UserId()
	// gen primary id
	order.SQLModel = common.GenNewModel()
	// gen order name
	orderName, err := sequence.Next(db, "S", 5, 1)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("cannot generate order name with sequence").WithDebug(err.Error())
	}
	order.Name = orderName
	// business end

	if err := db.Create(&order).Error; err != nil {
		return nil, core.ErrInternalServerError.WithError("cannot create order").WithDebug(err.Error())
	}
	return order, nil
}
