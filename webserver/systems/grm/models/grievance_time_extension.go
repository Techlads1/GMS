package models

import "time"

type GrievanceTimeExtension struct {
	Id              int       `json:"id,omitempty" param:"id" form:"id" validate:"omitempty,numeric"`
	GrievanceId 	int    	  `json:"grievance_id" form:"grievance_id"`
	GFUId 			int       `json:"gfu_id" form:"gfu_id"`
	State       	string    `json:"state" form:"state" validate:"required"`
	Description 	string    `json:"description" form:"description" validate:"required"`
	Comment       	string    `json:"comment" form:"comment" validate:"required"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
}