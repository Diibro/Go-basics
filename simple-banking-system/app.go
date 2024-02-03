package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to the Go Bank")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Print("Your choice: ")
	var choice int
	fmt.Scan(&choice)
	if choice == 1 {
		fmt.Println("Your account balance:", accountBalance)
	} else if choice == 2 {
		var newAmount float64
		fmt.Print("Deposit Amount: ")
		fmt.Scan(&newAmount)
		if newAmount <= 0 {
			fmt.Println("Invalid deposit amount. Amount must be greater than 0!")
		} else {
			accountBalance += newAmount
			fmt.Println("Deposit successful!! New Balance:", accountBalance)
		}

	} else if choice == 3 {
		var withdrawAmount float64
		fmt.Print("Withdraw Amount: ")
		fmt.Scan(&withdrawAmount)
		if withdrawAmount <= 0 {
			fmt.Println("Invalid Withdraw amount. Amount must be greater than 0!")
		} else if withdrawAmount > accountBalance {
			fmt.Println("Insufficient  funds in your account.")
		} else {
			accountBalance -= withdrawAmount
			fmt.Println("Withdraw successful!! New Balance:", accountBalance)
		}

	}
}
