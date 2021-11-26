package models

import "time"

type GrievanceAppeal struct {
	Id                    		int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	ReferenceNumber       		string    `json:"reference_number" form:"reference_number"`
	GrievanceReferenceNumber  string    `json:"grievance_reference_number" form:"grievance_reference_number"`
	DesiredOutcome       			string    `json:"comment" form:"comment" validate:"required"`
	Description 							string    `json:"description" form:"description" validate:"required"`
	GrievantId             		int       `json:"grievant_id,omitempty" form:"grievant_id" validate:"omitempty,numeric"`
	UpdatedAt             		time.Time `json:"updated_at,omitempty"`
	CreatedAt             		time.Time `json:"created_at,omitempty"`
}