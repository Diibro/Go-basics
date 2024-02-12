package main

import (
	"fmt"

	"dushime.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const BalanceFile string = "balance.txt"

func main() {

	fmt.Println("Welcome to the Go Bank")
	fmt.Println("Reach us 24/7 ", randomdata.PhoneNumber())
	accountBalance, err := fileops.ReadBalance(BalanceFile)

	if err != nil {
		fmt.Println("ERROR------")
		fmt.Println(err)
		fmt.Println("------------")
	}

	for {
		presentOptions()
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Your account balance:", accountBalance)
		case 2:
			var newAmount float64
			fmt.Print("Deposit Amount: ")
			fmt.Scan(&newAmount)
			if newAmount <= 0 {
				fmt.Println("Invalid deposit amount. Amount must be greater than 0!")
				continue
			} else {
				accountBalance += newAmount
				fileops.SaveBalance(accountBalance, BalanceFile)
				fmt.Println("Deposit successful!! New Balance:", accountBalance)
			}
		case 3:
			var withdrawAmount float64
			fmt.Print("Withdraw Amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid Withdraw amount. Amount must be greater than 0!")
			} else if withdrawAmount > accountBalance {
				fmt.Println("Insufficient  funds in your account.")
			} else {
				accountBalance -= withdrawAmount
				fileops.SaveBalance(accountBalance, BalanceFile)
				fmt.Println("Withdraw successful!! New Balance:", accountBalance)
			}
		case 4:
			fmt.Println("Byee")
			fmt.Println("Thank you for using our bank")
			return
		default:
			fmt.Println("Invalid choice!!!")
		}
	}

}
