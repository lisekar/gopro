package main

import (
	"fmt"
)

func main() {
	accountBalance := 10000.0
	fmt.Print("Welcome to Go Bank ! \n")
	fmt.Print("What do you want to do ? \n")
	fmt.Print("1. Create Account \n2. Deposit Amount \n3. Withdraw Amount \n4. Check Balance \n5. Exit \n")

	fmt.Print("Enter your choice (1-5) : ")
	var choice int
	fmt.Scan(&choice)

	fmt.Printf("You have selected option : %d\n", choice)

	switch choice {
	case 1:
		fmt.Print("Creating account...\n")
	case 2:
		fmt.Print("Enter amount to deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Printf("Deposited %.2f. New balance: %.2f\n", depositAmount, accountBalance)
	case 3:
		fmt.Print("Enter amount to withdraw: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		if withdrawAmount > accountBalance {
			fmt.Print("Insufficient funds!\n")
		} else {
			accountBalance -= withdrawAmount
			fmt.Printf("Withdrew %.2f. New balance: %.2f\n", withdrawAmount, accountBalance)
		}
	case 4:
		fmt.Printf("Current balance: %.2f\n", accountBalance)
	case 5:
		fmt.Print("Exiting...\n")
	default:
		fmt.Print("Invalid option!\n")
	}
	fmt.Print("Thank you for using Go Bank!\n")
}
