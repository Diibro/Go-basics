package main

import (
	"fmt"

	"dushime.com/user"
)

func main() {
	fmt.Println("Studying about structs")
	firstName := getUserData("Enter your first name: ")
	secondName := getUserData("Enter your last name: ")
	birthDate := getUserData("Enter your birthDate: ")

	userData, err := user.NewUser(firstName, secondName, birthDate)
	if err != nil {
		fmt.Println(err)
	} else {
		userData.ClearName()
		userData.OutPutInfo()
	}

}

func getUserData(prompText string) string {
	fmt.Print(prompText)
	value := ""
	fmt.Scanln(&value)
	return value
}
