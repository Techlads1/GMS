package models

import "time"

type GrievanceState struct {
	Id                    int       `json:"id,omitempty" param:"id" form:"id" validate:"omitempty,numeric"`
	Name       						string    `json:"name" form:"name" validate:"required"`
	CodeName       				string    `json:"code_name" form:"code_name" validate:"required"`
	Description 					string    `json:"description" form:"description" validate:"required"`
	Days                  int       `json:"days,omitempty" form:"days" validate:"omitempty,numeric"`
	UpdatedAt             time.Time `json:"updated_at,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}