package main

import "fmt"

func main() {

	fmt.Println("This is a profit calculator")
	revenue := userInput("Enter the revenue: ")
	expenses := userInput("Enter the expenses: ")
	taxRate := userInput("Enter the tax rate: ")

	ebtAmount, profit, ratio := calcEPR(revenue, expenses, taxRate)
	formattedEV := fmt.Sprintf("The earnings before tax: %v", ebtAmount)
	fmt.Print(formattedEV)
	fmt.Println("\nThe Profit made:  ", profit)
	fmt.Println("The ratio is: ", fmt.Sprintf("%.2f", ratio))
}

func userInput(text string) (value float64) {
	fmt.Print(text)
	fmt.Scan(&value)
	return
}

func calcEPR(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebtAmount := revenue - expenses
	profit := ebtAmount - (1 - taxRate/100)
	ratio := ebtAmount / profit
	return ebtAmount, profit, ratio
}
