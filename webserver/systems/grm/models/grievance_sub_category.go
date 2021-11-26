package models

import "time"

type GrievanceSubCategory struct {
	Id                    int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name       						string    `json:"name" form:"name" validate:"required"`
	CodeName       				string    `json:"code_name" form:"code_name" validate:"required"`
	Description 					string    `json:"description" form:"description" validate:"required"`
	UpdatedAt             time.Time `json:"updated_at,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}