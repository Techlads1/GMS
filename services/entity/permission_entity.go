package entity

import (
	"errors"
	"time"
)

var ErrInvalidPermissionEntity = errors.New("invalid permission")

// Permission Struct
type Permission struct {
	ID         int32
	Path       string
	Method     string
	Service    string
	SubService string
	Action     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

// NewPermission
func NewPermission(path, method, service, subSubservice, action string) (*Permission, error) {
	Permission := &Permission{
		Path:       path,
		Method:     method,
		Service:    service,
		SubService: subSubservice,
		Action:     action,
	}
	err := Permission.Validate()
	if err != nil {
		return nil, err
	}
	return Permission, nil

}

// Validate Permission data
func (p *Permission) Validate() error {
	if p.Path == "" || p.Method == "" || p.Service == "" || p.SubService == "" || p.Action == "" {
		return ErrInvalidPermissionEntity
	}
	return nil
}
