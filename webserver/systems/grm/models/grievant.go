package models

import "time"

type Grievant struct {
	Id                    int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	UserID       					int    		`json:"user_id" form:"user_id" validate:"required"`
	UpdatedAt             time.Time `json:"updated_at,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}
