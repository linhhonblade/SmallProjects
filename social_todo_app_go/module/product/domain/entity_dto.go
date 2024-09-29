package productdomain

import (
	"github.com/google/uuid"
	"social_todo_app_go/common"
)

type ProductCreateDTO struct {
	common.SQLModel
	Name        string        `json:"name" gorm:"column:name;"`
	Type        string        `json:"type" gorm:"column:type;"`
	Description string        `json:"description" gorm:"column:description;"`
	CategoryId  uuid.UUID     `json:"category_id" gorm:"column:category_id;"`
	Image       *common.Image `json:"image" gorm:"column:image;"`
}

func (ProductCreateDTO) TableName() string { return Product{}.TableName() }

type ProductUpdateDTO struct {
	Name        *string       `json:"name" gorm:"column:name;"`
	Description *string       `json:"description" gorm:"column:description;"`
	Image       *common.Image `json:"image" gorm:"column:image;"`
	Type        *string       `json:"type" gorm:"column:type;"`
	CategoryId  *uuid.UUID    `json:"category_id" gorm:"column:category_id;"`
}

func (ProductUpdateDTO) TableName() string { return Product{}.TableName() }
