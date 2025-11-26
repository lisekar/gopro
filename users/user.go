package user

import (
	"errors"
	"time"
)

// defining a struct - custom type
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func NewUser(firstName, lastName, birthDate string) User {
	return User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
}

func NewUserPointer(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("input should not be empty")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
