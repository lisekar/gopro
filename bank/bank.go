package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/livisekar/gopro/fileops"
)

func main() {
	// accountBalance := 10000.0
	var accountBalance, err = fileops.GetBalanceFromFile(fileops.AccountBalanceFile) // get balance from file
	if err != nil {
		fmt.Println("ERROR: ", err)
		fmt.Print("------------------------------ \n")
		// panic("Can't coontinue, sorry ! \n")
	}
	fmt.Print("Welcome to Go Bank ! \n")
	fmt.Print(randomdata.SillyName() + "\n")
	for {
		fmt.Print("What do you want to do ? \n")
		fmt.Print("1. Create Account \n2. Deposit Amount \n3. Withdraw Amount \n4. Check Balance \n5. Exit \n")
		// presentOptions()

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
			if depositAmount <= 0 {
				fmt.Print("Deposit amount must be greater than 0 \n")
				return
			}
			accountBalance += depositAmount
			fmt.Printf("Deposited %.2f. New balance: %.2f\n", depositAmount, accountBalance)
			fileops.WriteBalanceToFile(accountBalance, fileops.AccountBalanceFile)
			continue
		case 3:
			fmt.Print("Enter amount to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Print("Insufficient funds!\n")
				continue
			} else if withdrawAmount <= 0 {
				fmt.Print("Withdraw amount must be greater than 0 \n")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Printf("Withdrew %.2f. New balance: %.2f\n", withdrawAmount, accountBalance)
			fileops.WriteBalanceToFile(accountBalance, fileops.AccountBalanceFile)
		case 4:
			fmt.Printf("Current balance: %.2f\n", accountBalance)
			continue
		case 5:
			fmt.Print("Exiting... Goodbye!\n")
			return
		default:
			fmt.Print("Invalid option!\n")
			return
		}
		fmt.Print("Thank you for using Go Bank!\n")
	}

}
