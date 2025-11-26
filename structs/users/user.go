package user

import (
	"errors"
	"fmt"
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

func (u User) PrintDetails() {
	fmt.Println("User Details:")
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birth Date:", u.birthDate)
	fmt.Println("Account Created At:", u.createdAt.Format(time.RFC1123))
}

func (u *User) PrintDetailsPointer() {
	fmt.Println("User Details:")
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birth Date:", u.birthDate)
	fmt.Println("Account Created At:", u.createdAt.Format(time.RFC1123))
}

func PrintUserDetails(user User) {
	fmt.Println("User Details:")
	fmt.Println("First Name:", user.firstName)
	fmt.Println("Last Name:", user.lastName)
	fmt.Println("Birth Date:", user.birthDate)
	fmt.Println("Account Created At:", user.createdAt.Format(time.RFC1123))
}

func StructPointer(u *User) {
	fmt.Println("Accessing User First Name via Pointer:", (*u).firstName)
	fmt.Println("Accessing User Last Name via Pointer:", (*u).lastName)
	fmt.Println("Accessing User Birth Date via Pointer:", (*u).birthDate)
	fmt.Println("Accessing User Created At via Pointer:", (*u).createdAt.Format(time.RFC1123))
}

func (u *User) ClearUserName() { // add method into struct with pointer
	u.firstName = ""
	u.lastName = ""
	u.birthDate = ""
	fmt.Println("Cleared User Details inside clearUserName method:")
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birth Date:", u.birthDate)
}
