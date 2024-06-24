package main

import (
	"fmt"

	"example.com/structs-go/user"
)

type str string

func (text str) log() {
	fmt.Println(text)
}

func main() {
	appUser, err := user.New(getUserData("Please enter your first name: "), getUserData("Please enter your last name: "), getUserData("Please enter your birthdate (MM/DD/YYYY): "))

	if err != nil {
		fmt.Println(err)
	}

	appUser.Output()
	appUser.ClearName()
	appUser.Output()

	admin := user.NewAdmin("perosamurilo@gmail.com", "Murilo")

	admin.Output()
	admin.ClearName()
	admin.Output()

	var name str = "Murilo"
	name.log()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
