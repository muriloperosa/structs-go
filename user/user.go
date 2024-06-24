package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// Struct Embedding
type Admin struct {
	email    string
	password string
	User
}

func (u User) Output() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *User) ClearName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName string, lastName string, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birth date are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email string, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "System",
			lastName:  "Admin",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}
