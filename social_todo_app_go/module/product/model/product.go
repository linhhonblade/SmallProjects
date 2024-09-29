package model

import (
	"errors"
	"github.com/google/uuid"
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

type ProductCreation struct {
	common.SQLModel
	Name        string        `json:"name" gorm:"column:name;"`
	Type        string        `json:"type" gorm:"column:type;"`
	Description string        `json:"description" gorm:"column:description;"`
	CategoryId  uuid.UUID     `json:"category_id" gorm:"column:category_id;"`
	Image       *common.Image `json:"image" gorm:"column:image;"`
}

func (ProductCreation) TableName() string { return Product{}.TableName() }

type ProductUpdate struct {
	Name        string        `json:"name" gorm:"column:name;"`
	Description *string       `json:"description" gorm:"column:description;"`
	Status      string        `json:"status" gorm:"column:status;"`
	Image       *common.Image `json:"image" gorm:"column:image;"`
	Type        string        `json:"type" gorm:"column:type;"`
}

func (ProductUpdate) TableName() string { return Product{}.TableName() }