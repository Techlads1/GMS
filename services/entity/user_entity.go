package entity

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//User defines class user data
type User struct {
	ID              int32
	Name            string
	Password        string
	Email           string
	Token           string
	LoginAttempt    int32
	Active          bool
	UserCategoryID  int32
	Roles           []int32
	EmailVerifiedAt time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

//NewUser creates a new status
func NewUser(name string, password string, email string, userCategoryID int32) (*User, error) {

	user := &User{
		Name:           name,
		Password:       password,
		Email:          email,
		UserCategoryID: userCategoryID,
	}

	pwd, err := generatePassword(password)
	if err != nil {
		return nil, err
	}

	user.Password = pwd
	err = user.Validate()
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Validate user data
func (u *User) Validate() error {
	if u.Name == "" || u.Password == "" || u.Email == "" || u.UserCategoryID < 1 {
		return errors.New("error validating user")
	}
	return nil
}

// Validate user data
func (u *User) UpdateUserValidate() error {
	if u.Name == "" || u.Email == "" || u.ID < 1 {
		return errors.New("invalid user name or password")
	}
	return nil
}

//ValidatePassword validate user password
func (u *User) ValidatePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err != nil {
		return err
	}
	return nil
}

func generatePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
