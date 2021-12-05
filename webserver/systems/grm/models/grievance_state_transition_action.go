package models

import "time"

type GrievanceStateTransitionAction struct {
	Id                    int       `json:"id,omitempty" param:"id" form:"id" validate:"omitempty,numeric"`
	GrievantId       			int    		`json:"grievant_id" form:"grievant_id" validate:"required"`
	Data       						string    `json:"code_name" form:"code_name" validate:"required"`
	GrievanceStateId      int    		`json:"grievant_state_id" form:"grievance_state_id" validate:"required"`
	Description 					string    `json:"description" form:"description" validate:"required"`
	UpdatedAt             time.Time `json:"updated_at,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}