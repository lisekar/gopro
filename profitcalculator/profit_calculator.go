package main

import (
	"fmt"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxrate float64

	// fmt.Print("Enter your Profit amount : ")
	// fmt.Scan(&revenue)

	revenue := getUserInput("Revenue : ") // getting user input using function

	// fmt.Print("Enter your Profit Rate : ")
	// fmt.Scan(&expenses)

	expenses := getUserInput("Expenses : ")

	// fmt.Print("Enter the taxrate : ")
	// fmt.Scan(&taxrate)

	taxrate := getUserInput("Tax Rate (in %) : ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxrate)

	fmt.Printf("Earnings Before Tax : %.1f	\n", ebt)
	fmt.Printf("Profit After Tax : %.1f	\n", profit)
	fmt.Printf("Profit Ratio : %.1f%%\n", ratio)
}

func calculateFinancials(revenue, expenses, taxrate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt - (ebt * taxrate / 100)
	if profit != 0 {
		ratio = (ebt / profit) * 100
	}
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}
