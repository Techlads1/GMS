package models

import "time"

type Grievance struct {
	Id                    	int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name       							string    `json:"name" form:"name" validate:"required"`
	ReferenceNumber       	string    `json:"reference_number" form:"reference_number"`
	Comment       					string    `json:"comment" form:"comment" validate:"required"`
	State       						string    `json:"state" form:"state" validate:"required"`
	Description 						string    `json:"description" form:"description" validate:"required"`
	LocationOccurred      	string    `json:"location_occurred" form:"location_occurred" validate:"required"`
	FillingModeId         	int       `json:"filling_mode_id,omitempty" form:"filling_mode_id" validate:"omitempty,numeric"`
	GrievanceSubCategoryId	int       `json:"grievance_sub_category_id,omitempty" form:"grievance_sub_category_id" validate:"omitempty,numeric"`
	GrievantId              int       `json:"grievant_id,omitempty" form:"grievant_id" validate:"omitempty,numeric"`
	GrievantGroupId         int       `json:"grievant_group_id,omitempty" form:"grievant_group_id" validate:"omitempty,numeric"`
	UpdatedAt            		time.Time `json:"updated_at,omitempty"`
	CreatedAt             	time.Time `json:"created_at,omitempty"`
}