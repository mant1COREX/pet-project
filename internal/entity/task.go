package entity

import "time"

type Task struct {
	Id          int       `json:"id" validate:"omitempty,gte=1" db:"id"`
	Title       string    `json:"title" validate:"required,min=1" db:"title"`
	Description string    `json:"description" validate:"omitempty,min=1" db:"description"`
	Status      string    `json:"status" validate:"omitempty,oneof=new in_progress done" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
