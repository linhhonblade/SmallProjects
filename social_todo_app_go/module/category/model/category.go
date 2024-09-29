package model

import (
	"errors"
	"social_todo_app_go/common"
	"strings"
)

var (
	ErrNameCannotEmpty = errors.New("Title cannot be empty")
	ErrCategoryDeleted = errors.New("Item is deleted")
)

const (
	EntityName = "Category"
)

type Category struct {
	common.SQLModel
	Name        string        `json:"name" gorm:"column:name;"`
	Description string        `json:"description" gorm:"column:description;"`
	Image       *common.Image `json:"image" gorm:"column:image;"`
}

func (Category) TableName() string { return "category" }

type CategoryCreation struct {
	common.SQLModel
	Name        string        `json:"name" gorm:"column:name;"`
	Description string        `json:"description" gorm:"column:description;"`
	Image       *common.Image `json:"image" gorm:"column:image;"`
}

func (CategoryCreation) TableName() string { return Category{}.TableName() }

func (i *CategoryCreation) Validate() error {
	i.Name = strings.TrimSpace(i.Name)

	if i.Name == "" {
		return ErrNameCannotEmpty
	}
	return nil
}

type CategoryUpdate struct {
	Name        string        `json:"name" gorm:"column:name;"`
	Description *string       `json:"description" gorm:"column:description;"`
	Status      string        `json:"status" gorm:"column:status;"`
	Image       *common.Image `json:"image" gorm:"column:image;"`
}

func (CategoryUpdate) TableName() string { return Category{}.TableName() }
