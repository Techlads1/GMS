package models

import "time"

type GrievanceStateAction struct {
	Id                    int       `json:"id,omitempty" param:"id" form:"id" validate:"omitempty,numeric"`
	Name       						string    `json:"name" form:"name" validate:"required"`
	RolePerformAction     string    `json:"role_perform_action" form:"role_perform_action" validate:"required"`
	StateId               int       `json:"state_id,omitempty" form:"state_id" validate:"omitempty,numeric"`
	UpdatedAt             time.Time `json:"updated_at,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}