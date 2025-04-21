package usecase

import (
	"context"
	"github.com/google/uuid"
	"social_todo_app_go/module/order/domain"
)

type UseCase interface {
	AddProductToOrder(ctx context.Context, productIDs []uuid.UUID, orderID uuid.UUID) error
	GetDraftOrder(ctx context.Context) (*domain.Order, error)
	CreateDraftOrder(ctx context.Context, dto OrderDTO) (*domain.Order, error)
}

type OrderRepository interface {
	OrderQueryRepository
	OrderCommandRepository
}

type OrderQueryRepository interface {
	FindDraftOrderByUserId(ctx context.Context, userID uuid.UUID) (*OrderDTO, error)
}

type OrderCommandRepository interface {
	Update(ctx context.Context, order *OrderDTO) error
	Create(ctx context.Context, order *OrderDTO) error
}
