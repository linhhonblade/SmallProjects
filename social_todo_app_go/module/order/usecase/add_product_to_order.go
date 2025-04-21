package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/linhhonblade/service-context/core"
	"social_todo_app_go/common"
	"social_todo_app_go/module/user/domain"
)

type addProductToOrderUC struct {
	orderRepo OrderRepository
}

func NewAddProductToOrderUC(orderRepo OrderRepository) *addProductToOrderUC {
	return &addProductToOrderUC{orderRepo: orderRepo}
}

func (uc *addProductToOrderUC) AddProductToOrder(ctx context.Context, dto AddProductToOrderDTO) error {
	// 1. Retrieve user id from Requester
	userId := dto.Requester.UserId()

	// 1. Find draft order by user id
	if dto.OrderID == nil {

	}
	orderEntity, err := uc.orderRepo.FindDraftOrderByUserId(ctx, userId)
	if err != nil {
		return core.ErrBadRequest.WithDebug(err.Error())
	}
	if orderEntity == nil {
		// Create new draft order
		panic("To be implemented")
	}

	// 2. Create new order line
	orderLineEntity, err := domain.NewOrderLine(common.GenUUID(), dto.OrderID, dto.ProductID, dto.Quantity)

	return nil
}
