package model

import (
	"errors"
	"server/common"
)

var (
	ErrTitleIsBlank = errors.New("title can't be blank")
)

type TodoItem struct {
	common.SQLModel
	Title       string      `json:"title" gorm:"column:title;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

type TodoItemCreate struct {
	Id          int         `json:"-" gorm:"column:id;"`
	Title       string      `json:"title" gorm:"column:title;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
	Description string      `json:"description" gorm:"column:description;"`
}

func (TodoItemCreate) TableName() string {
	return TodoItem{}.TableName()
}

type TodoItemUpdate struct {
	Title       *string     `json:"title" gorm:"column:title;"`
	Description *string     `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}
