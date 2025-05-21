package query

import (
	"github.com/google/uuid"
	"social_todo_app_go/common"
	"time"
)

type OrderDTO struct {
	common.SQLModel
	Name       string         `json:"name" gorm:"column:name;"`
	UserId     uuid.UUID      `json:"user_id" gorm:"column:user_id;"`
	DateOrder  *time.Time     `json:"date_order,omitempty" gorm:"column:date_order;"`
	OrderItems []OrderItemDTO `json:"order_items" gorm:"foreignKey:OrderId"`
}

func (OrderDTO) TableName() string {
	return "order"
}

type OrderItemDTO struct {
	common.SQLModel
	OrderId    uuid.UUID   `json:"order_id" gorm:"column:order_id;" binding:"required"`
	ProductId  uuid.UUID   `json:"product_id" gorm:"column:product_id;"`
	Product    *ProductDTO `json:"product" gorm:"foreignKey:ProductId"`
	Quantity   int         `json:"quantity" gorm:"column:quantity;"`
	UnitPrice  float64     `json:"unit_price" gorm:"column:unit_price;"`
	TotalPrice float64     `json:"total_price" gorm:"column:total_price;"`
}

func (OrderItemDTO) TableName() string {
	return "order_item"
}

type ProductDTO struct {
	Id   uuid.UUID `json:"id" gorm:"column:id;"`
	Name string    `json:"name" gorm:"column:name;"`
}

func (ProductDTO) TableName() string {
	return "product"
}
