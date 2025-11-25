package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

// function to write balance to a file
func writeBalanceToFile(balance float64, filename string) {
	balanceText := fmt.Sprintf("Current Balance: %.2f\n", balance) // format balance text
	os.WriteFile(filename, []byte(balanceText), 0644)              // write to file with appropriate permissions
}

func getBalanceFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 10000, errors.New("fail to find balance file") // return 0 if file does not exist or error occurs
	}
	balanceText := string(data)                                   // float64 cant be directly converted from []byte
	balanceValue, err := strconv.ParseFloat(balanceText[17:], 64) // extract balance value from text
	if err != nil {
		return 10000, errors.New("fail to parse balance value") // return 0 if parsing fails
	}
	return balanceValue, nil
}

func main() {
	// accountBalance := 10000.0
	var accountBalance, err = getBalanceFromFile(accountBalanceFile) // get balance from file
	if err != nil {
		fmt.Println("ERROR: ", err)
		fmt.Print("------------------------------ \n")
		// panic("Can't coontinue, sorry ! \n")
	}
	fmt.Print("Welcome to Go Bank ! \n")
	for {
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
			if depositAmount <= 0 {
				fmt.Print("Deposit amount must be greater than 0 \n")
				return
			}
			accountBalance += depositAmount
			fmt.Printf("Deposited %.2f. New balance: %.2f\n", depositAmount, accountBalance)
			writeBalanceToFile(accountBalance, accountBalanceFile)
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
			writeBalanceToFile(accountBalance, accountBalanceFile)
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
