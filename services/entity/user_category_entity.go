package entity

import (
	"errors"
	"time"
)

var ErrInvalidUserCategoryEntity = errors.New("invalid user category")

// UserCategory DatasStructure
type UserCategory struct {
	ID        int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	CreatedBy int32
}

// NewUserCategory
func NewUserCategory(name string, id int32) (*UserCategory, error) {
	UserCategory := &UserCategory{
		Name:      name,
		CreatedBy: id,
	}
	err := UserCategory.Validate()
	if err != nil {
		return nil, err
	}
	return UserCategory, nil

}

// Validate UserCategory
func (st *UserCategory) Validate() error {
	if st.Name == "" || st.CreatedBy < 1 {
		return ErrInvalidUserCategoryEntity
	}
	return nil
}
