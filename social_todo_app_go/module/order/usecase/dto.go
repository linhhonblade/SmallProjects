package usecase

import (
	"github.com/google/uuid"
	"social_todo_app_go/common"
)

type OrderDTO struct {
	Id         uuid.UUID      `gorm:"column:id" json:"id"`
	UserId     uuid.UUID      `gorm:"column:user_id" json:"user_id"`
	ExpireAt   string         `gorm:"column:expire_at" json:"expire_at"`
	Status     string         `gorm:"column:status" json:"status"`
	OrderLines []OrderLineDTO `json:"order_lines"`
}

type OrderLineDTO struct {
	Id        uuid.UUID   `gorm:"column:id" json:"id"`
	OrderId   uuid.UUID   `gorm:"column:order_id" json:"order_id"`
	ProductId uuid.UUID   `gorm:"column:product_id" json:"product_id"`
	Product   *ProductDTO `json:"product" gorm:"-"`
}

type ProductDTO struct {
	Id   uuid.UUID
	Name string `json:"name" gorm:"column:name;"`
}

type AddProductToOrderDTO struct {
	ProductIDs []uuid.UUID      `json:"product_ids"`
	OrderID    uuid.UUID        `json:"order_id"`
	Requester  common.Requester `json:"-"`
}

type CreateDraftOrderDTO struct {
	UserId uuid.UUID `gorm:"column:user_id" json:"user_id"`
}
