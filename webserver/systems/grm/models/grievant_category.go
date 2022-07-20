package models

import (
	"time"
)

type GrievantCategory struct {
	Id                    int       `json:"id,omitempty" param:"id" form:"id" validate:"omitempty,numeric"`
	Name       						string    `json:"name" form:"name" validate:"required"`
	Description 					string    `json:"description" form:"description" validate:"required"`
	UpdatedAt             time.Time `json:"updated_at,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}


type GrievantCategoryRepository interface {

	Create(arg GrievantCategory) (int, error)
	Get(id int) (GrievantCategory, error) 
	Update(arg *GrievantCategory) (int, error)
	All() ([]*GrievantCategory, error)
	Delete(id int) error

}

