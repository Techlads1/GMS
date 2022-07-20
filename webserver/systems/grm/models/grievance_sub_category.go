package models

import "time"

type GrievanceSubCategory struct {
	Id                    int       `json:"id,omitempty" param:"id" form:"id" validate:"omitempty,numeric"`
	Name       						string    `json:"name" form:"name" validate:"required"`
	CodeName       				string    `json:"code_name" form:"code_name" validate:"required"`
	Description 					string    `json:"description" form:"description" validate:"required"`
	GrievanceCategoryId   int				`json:"grievance_category_id" form:"grievance_category_id" validate:"required"`
	UpdatedAt             time.Time `json:"updated_at,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}