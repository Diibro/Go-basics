package main

import "fmt"

func main() {
	fmt.Println("Studying about go pointers")
	age := 32
	agePointer := &age
	fmt.Printf("Age: %v\n", age)
	getAdultAge(agePointer)
	fmt.Printf("Age: %v\n", age)
}

func getAdultAge(ageP *int) {
	*ageP = *ageP - 18
}
