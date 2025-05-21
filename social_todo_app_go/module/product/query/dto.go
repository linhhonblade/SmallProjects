package query

import (
	"github.com/google/uuid"
	"social_todo_app_go/common"
)

type ProductDTO struct {
	common.SQLModel
	Name       string       `json:"name" gorm:"column:name;"`
	Type       string       `json:"type" gorm:"column:type;"`
	CategoryId uuid.UUID    `json:"category_id" gorm:"column:category_id;"`
	Category   *CategoryDTO `json:"category" gorm:"-"`
}

type CategoryDTO struct {
	Id   uuid.UUID `json:"id" gorm:"column:id;"`
	Name string    `json:"name" gorm:"column:name;"`
}

func (CategoryDTO) TableName() string {
	return "category"
}
func (ProductDTO) TableName() string {
	return "product"
}
