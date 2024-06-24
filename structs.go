package main

import (
	"fmt"

	"example.com/structs-go/user"
)

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
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
