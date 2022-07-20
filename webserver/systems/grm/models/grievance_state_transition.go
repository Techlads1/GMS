package models

import "time"

type GrievanceStateTransition struct {
	Id                    int       `json:"id,omitempty" param:"id" form:"id" validate:"omitempty,numeric"`
	FromStateId      			int    		`json:"from_state_id" form:"from_state_id" validate:"required"`
	ToStateId      				int    		`json:"to_state_id" form:"to_state_id" validate:"required"`
	Description 					string    `json:"description" form:"description" validate:"required"`
	UpdatedAt             time.Time `json:"updated_at,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}