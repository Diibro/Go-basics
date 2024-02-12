package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	fmt.Println("Studying about structs")
	var userData user
	firstName := getUserData("Enter your first name: ")
	secondName := getUserData("Enter your last name: ")
	birthDate := getUserData("Enter your birthDate: ")

	userData = user{
		firstName: firstName,
		lastName:  secondName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}

	fmt.Println(userData.firstName, " ", userData.lastName, ", ", userData.createdAt)
}

func getUserData(prompText string) string {
	fmt.Print(prompText)
	value := ""
	fmt.Scan(&value)
	return value
}
