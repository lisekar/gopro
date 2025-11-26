package main

import (
	"fmt"

	user "github.com/livisekar/gopro/structs/users"
)

// // defining a struct - custom type
// type User struct {
// 	firstName string
// 	lastName  string
// 	birthDate string
// 	createdAt time.Time
// }

// // func newUser(firstName, lastName, birthDate string) User {
// // 	return User{
// // 		firstName: firstName,
// // 		lastName:  lastName,
// // 		birthDate: birthDate,
// // 		createdAt: time.Now(),
// // 	}
// // }

// func newUserPointer(firstName, lastName, birthDate string) (*User, error) {
// 	if firstName == "" || lastName == "" || birthDate == "" {
// 		return nil, errors.New("input should not be empty")
// 	}
// 	return &User{
// 		firstName: firstName,
// 		lastName:  lastName,
// 		birthDate: birthDate,
// 		createdAt: time.Now(),
// 	}, nil
// }

func main() {
	firstName := getUserData("First Name: ")
	lastName := getUserData("Last Name: ")
	birthDate := getUserData("Birth Date (YYYY-MM-DD): ")

	// appUser := User{
	// 	firstName: firstName,
	// 	lastName:  lastName,
	// 	birthDate: birthDate,
	// 	createdAt: time.Now(),
	// }
	var appUser *user.User

	// appUser, err := newUserPointer(firstName, lastName, birthDate) // access user structs before package creation

	// appUser = &user.User{
	// 	FirstName: "Livi",  		// access property from structs , make the fieldname first char as capital letter in package.
	// }

	appUser, err := user.NewUserPointer(firstName, lastName, birthDate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	// appUserPointer := newUserPointer(firstName, lastName, birthDate)

	// printUserDetails(appUser)
	// structPointer(&appUser)

	appUser.PrintDetails()        // calling method with value receiver
	appUser.PrintDetailsPointer() // calling method with pointer receiver

	appUser.ClearUserName() // calling method that modifies struct via pointer receiver
	appUser.PrintDetails()  // after clearing user details
}

// func (u User) printDetails() {
// 	fmt.Println("User Details:")
// 	fmt.Println("First Name:", u.firstName)
// 	fmt.Println("Last Name:", u.lastName)
// 	fmt.Println("Birth Date:", u.birthDate)
// 	fmt.Println("Account Created At:", u.createdAt.Format(time.RFC1123))
// }

// func (u *User) printDetailsPointer() {
// 	fmt.Println("User Details:")
// 	fmt.Println("First Name:", u.firstName)
// 	fmt.Println("Last Name:", u.lastName)
// 	fmt.Println("Birth Date:", u.birthDate)
// 	fmt.Println("Account Created At:", u.createdAt.Format(time.RFC1123))
// }

// func printUserDetails(user User) {
// 	fmt.Println("User Details:")
// 	fmt.Println("First Name:", user.firstName)
// 	fmt.Println("Last Name:", user.lastName)
// 	fmt.Println("Birth Date:", user.birthDate)
// 	fmt.Println("Account Created At:", user.createdAt.Format(time.RFC1123))
// }

// func structPointer(u *User) {
// 	fmt.Println("Accessing User First Name via Pointer:", (*u).firstName)
// 	fmt.Println("Accessing User Last Name via Pointer:", (*u).lastName)
// 	fmt.Println("Accessing User Birth Date via Pointer:", (*u).birthDate)
// 	fmt.Println("Accessing User Created At via Pointer:", (*u).createdAt.Format(time.RFC1123))
// }

func getUserData(prompt string) string {
	var input string
	println(prompt)
	fmt.Scanln(&input)
	return input
}

// func (u *User) clearUserName() { // add method into struct with pointer
// 	u.firstName = ""
// 	u.lastName = ""
// 	u.birthDate = ""
// 	fmt.Println("Cleared User Details inside clearUserName method:")
// 	fmt.Println("First Name:", u.firstName)
// 	fmt.Println("Last Name:", u.lastName)
// 	fmt.Println("Birth Date:", u.birthDate)
// }
