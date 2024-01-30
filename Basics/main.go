package main

import (
	"fmt"
)

func main() {
	// var (
	// 	a int  = 1
	// 	b rune = 'b'
	// )
	var name1 string = "Dushime"
	var name2 string = "Bother"
	var finalName = stringConcat(name1, name2)
	fmt.Println(stringConcat("Dushime", "Bother"))
	fmt.Println(finalName)
	printUserIn()
	addnum()
}

func stringConcat(name1 string, name2 string) string {
	return name1 + " " + name2
}

func printUserIn() {
	var userInput string
	fmt.Scan(&userInput)
	fmt.Println("You entered: ", userInput)
}
