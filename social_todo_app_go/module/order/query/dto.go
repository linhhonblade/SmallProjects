package query

import (
	"github.com/google/uuid"
	"social_todo_app_go/common"
	"time"
)

type OrderDTO struct {
	common.SQLModel
	Name      string     `json:"name" gorm:"column:name;"`
	UserId    uuid.UUID  `json:"user_id" gorm:"column:user_id;"`
	DateOrder *time.Time `json:"date_order,omitempty" gorm:"column:date_order;"`
}

func (OrderDTO) TableName() string {
	return "order"
}
