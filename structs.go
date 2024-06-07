package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u user) output() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *user) clearName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName string, lastName string, birthdate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}

func main() {
	appUser := newUser(getUserData("Please enter your first name: "), getUserData("Please enter your last name: "), getUserData("Please enter your birthdate (MM/DD/YYYY): "))

	appUser.output()
	appUser.clearName()
	appUser.output()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
