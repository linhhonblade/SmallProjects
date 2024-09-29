package productdomain

import (
	"errors"
	"social_todo_app_go/common"
)

var (
	ErrProductDeleted = errors.New("Product is deleted")
)

const (
	EntityName = "Product"
)

type Product struct {
	common.SQLModel
	Name        string        `json:"name" gorm:"column:name;"`
	Description string        `json:"description" gorm:"column:description;"`
	Image       *common.Image `json:"image" gorm:"column:image;"`
	Type        string        `json:"type" gorm:"column:type;"`
	CategoryId  int           `json:"category_id" gorm:"column:category_id;"`
}

func (Product) TableName() string { return "product" }
