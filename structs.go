package main

import (
	"errors"
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

func newUser(firstName string, lastName string, birthdate string) (*user, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birth date are required")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	appUser, err := newUser(getUserData("Please enter your first name: "), getUserData("Please enter your last name: "), getUserData("Please enter your birthdate (MM/DD/YYYY): "))

	if err != nil {
		panic(err)
	}

	appUser.output()
	appUser.clearName()
	appUser.output()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
