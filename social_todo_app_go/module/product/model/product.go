package model

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
	Status      string        `json:"status" gorm:"column:status;"`
	Image       *common.Image `json:"image" gorm:"column:image;"`
	Type        string        `json:"type" gorm:"column:type;"`
	CategId     int           `json:"categ_id" gorm:"column:categ_id;"`
}

func (Product) TableName() string { return "product" }

type ProductCreation struct {
	Id          int           `json:"id" gorm:"column:id;"`
	Name        string        `json:"name" gorm:"column:name;"`
	Description string        `json:"description" gorm:"column:description;"`
	CategId     int           `json:"categ_id" gorm:"column:categ_id;"`
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
