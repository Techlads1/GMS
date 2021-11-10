package models

import "time"

//Department DataStructure
type Department struct {
	Id                    int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	DepartmentTitle       string    `json:"department_title" form:"department_title" validate:"required"`
	DepartmentDescription string    `json:"department_description" form:"department_description" validate:"required"`
	DepartmentSize        int       `json:"department_size"  form:"department_size" validate:"numeric"`
	CampusId              int       `json:"campus_id"  form:"campus_id" validate:"numeric"`
	CreatedBy             int       `json:"created_by,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}

//DepartmentFull DataStructure
type DepartmentFull struct {
	Id                    int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	DepartmentTitle       string    `json:"department_title" form:"department_title" validate:"required"`
	DepartmentDescription string    `json:"department_description" form:"department_description" validate:"required"`
	DepartmentSize        int       `json:"department_size"  form:"department_size" validate:"numeric"`
	CampusTitle           string    `json:"campus_id"  form:"campus_id" validate:"required"`
	CreatedBy             int       `json:"created_by,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}
