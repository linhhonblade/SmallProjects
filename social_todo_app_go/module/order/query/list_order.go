package query

import (
	"context"
	sctx "github.com/linhhonblade/service-context"
	"social_todo_app_go/common"
)

type listOrderQuery struct {
	sctx      sctx.ServiceContext
	requester common.Requester
}

func NewListOrderQuery(sctx sctx.ServiceContext, requester common.Requester) *listOrderQuery {
	return &listOrderQuery{sctx: sctx, requester: requester}
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
	if err := db.Offset(offset).Limit(param.Limit).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
