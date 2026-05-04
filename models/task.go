package models

import (
	"gorm.io/gorm"
)

type Status string

const (
	StatusPending    Status = "pending"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)

type Task struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	Status      Status `json:"status" gorm:"default:pending"`
}

type CreateTaskInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UpdateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}

type PatchTaskInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}
