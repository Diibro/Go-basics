package main

import "fmt"

func main() {
	var (
		revenue  float64
		expenses float64
		taxRate  float64
	)

	fmt.Println("This is a profit calculator")
	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)

	ebtAmount := revenue - expenses
	profit := ebtAmount - (ebtAmount * taxRate / 100)
	ratio := ebtAmount / profit

	fmt.Println("The earnings before tax: ", ebtAmount)
	fmt.Println("The Profit made:  ", profit)
	fmt.Println("The ratio is: ", ratio)

}
