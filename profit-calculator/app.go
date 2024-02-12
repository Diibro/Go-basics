package main

import (
	"errors"
	"fmt"
	"os"
)

const HistoryFile = "history.txt"

func main() {

	fmt.Println("This is a profit calculator")
	revenue, err := checkInput(userInput("Enter the revenue: "))
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		panic("Pram crushed..")
	}

	expenses, err := checkInput(userInput("Enter the expenses: "))
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		panic("Pram crushed..")
	}
	taxRate, err := checkInput(userInput("Enter the tax rate: "))

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		panic("Pram crushed..")
	}

	ebtAmount, profit, ratio := calcEPR(revenue, expenses, taxRate)
	saveCalc(ebtAmount, profit, ratio)
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

func checkInput(input float64) (float64, error) {
	if input < 0 {
		return 0, errors.New("Invalid input. Number must be greater than 0")
	}
	return input, nil
}

func saveCalc(ebtAmount, profit, ratio float64) {
	historyStr := fmt.Sprintf("%v,%v,%v", ebtAmount, profit, ratio)
	os.WriteFile(HistoryFile, []byte(historyStr), 0644)
}
