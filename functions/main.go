package main

import "fmt"

func main() {
	fmt.Println("Understanding functions")
	fmt.Println("Enter two numbers: ")
	var num1, num2 int
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	summed, subtracted := calc(num1, num2)
	fmt.Printf("The sum of the numbers: %v \nAnd the diff between the numbers is: %v", summed, subtracted)
}

func calc(num1, num2 int) (int, int) {
	added := num1 + num2
	subtracted := num2 - num1
	return added, subtracted

}
