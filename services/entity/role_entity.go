package entity

import (
	"errors"
	"time"
)

//Role DataStructure
type Role struct {
	ID          ID        `json:"id,omitempty"`
	Permissions []ID      `json:"permission_ids,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedBy   int32     `json:"created_by"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}

//NewRole
func NewRole(name, description string) (*Role, error) {
	role := &Role{
		Name:        name,
		Description: description,
	}
	err := role.Validate()
	if err != nil {
		return nil, err
	}
	return role, nil

}

func (r *Role) ValidateUpdateRole() error {
	if r.Name == "" {
		return errors.New("error validating role")
	}
	return nil
}

func (r *Role) Validate() error {
	if r.Name == "" {
		return errors.New("invalid role name")
	}
	return nil
}
