package models

import "time"

type GrievantGroup struct {
	Id                    int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name       						string    `json:"name" form:"name" validate:"required,max=200"`
	Description 					string    `json:"description" form:"description" validate:"required"`
	GrievantCategoryId 		int				`json:"grievant_category_id" form:"grievant_category_id" validate:"required"`
	UpdatedAt             time.Time `json:"updated_at,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}