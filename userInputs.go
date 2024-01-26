package main

import "fmt"

func addnum() {
	fmt.Print("Enter num 1: ")
	var num1 int
	fmt.Scanf("%d", &num1)
	fmt.Print("Enter num 1: ")
	var num2 int
	fmt.Scanf("%d", &num2)
	sum := num1 + num2
	fmt.Println("the sum of the enter numbers is: ", sum)
}
