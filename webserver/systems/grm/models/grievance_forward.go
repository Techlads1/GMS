package models

import "time"

type GrievanceForward struct {
	Id                  int       `json:"id,omitempty" param:"id" form:"id" validate:"omitempty,numeric"`
	GrievanceId       	int    	  `json:"grievance_id" form:"grievance_id"`
	State       		string    `json:"state" form:"state" validate:"required"`
	FromGFUId			int	  	  `json:"fromgfu_id" form:"fromgfu_id" validate:"required"`
	ToGFUId				int	  	  `json:"togfu_id" form:"togfu_id" validate:"required"`
	Description 		string    `json:"description" form:"description" validate:"required"`
	Comment       		string    `json:"comment" form:"comment" validate:"required"`
	UpdatedAt           time.Time `json:"updated_at,omitempty"`
	CreatedAt           time.Time `json:"created_at,omitempty"`
}