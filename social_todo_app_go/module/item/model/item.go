package model

import (
	"errors"
	"social_todo_app_go/common"
	"strings"
)

var (
	ErrTitleCannotEmpty = errors.New("Title cannot be empty")
	ErrItemDeleted      = errors.New("Item is deleted")
)

const (
	EntityName = "TodoItem"
)

type TodoItem struct {
	common.SQLModel
	Title       string        `json:"title" gorm:"column:title;"`
	Description string        `json:"description" gorm:"column:description;"`
	Status      string        `json:"status" gorm:"column:status;"`
	Image       *common.Image `json:"image" gorm:"column:image;"`
}

func (TodoItem) TableName() string { return "item" }

type TodoItemCreation struct {
	Id          int           `json:"id" gorm:"column:id;"`
	Title       string        `json:"title" gorm:"column:title;"`
	Description string        `json:"description" gorm:"column:description;"`
	Image       *common.Image `json:"image" gorm:"column:image;"`
}

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

func (i *TodoItemCreation) Validate() error {
	i.Title = strings.TrimSpace(i.Title)

	if i.Title == "" {
		return ErrTitleCannotEmpty
	}
	return nil
}

type TodoItemUpdate struct {
	Title       string  `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:description;"`
	Status      string  `json:"status" gorm:"column:status;"`
}

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }
