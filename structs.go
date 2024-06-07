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

func main() {
	appUser := user{
		firstName: getUserData("Please enter your first name: "),
		lastName:  getUserData("Please enter your last name: "),
		birthdate: getUserData("Please enter your birthdate (MM/DD/YYYY): "),
		createdAt: time.Now(),
	}

	outputUserData(&appUser)
}

func outputUserData(u *user) {
	// fmt.Println((*u).firstName, (*u).lastName, (*u).birthdate, (*u).createdAt)
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
