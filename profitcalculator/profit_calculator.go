package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxrate float64

	// fmt.Print("Enter your Profit amount : ")
	// fmt.Scan(&revenue)

	revenue, err1 := getUserInput("Revenue : ") // getting user input using function
	// print(err1)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Print("Enter your Profit Rate : ")
	// fmt.Scan(&expenses)

	expenses, err2 := getUserInput("Expenses : ")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Print("Enter the taxrate : ")
	// fmt.Scan(&taxrate)

	taxrate, err3 := getUserInput("Tax Rate (in %) : ")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Error: Invalid input(s) !")
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxrate)
	writeResultsInFile(resultsFile, ebt, profit, ratio)

	fmt.Printf("Earnings Before Tax : %.1f	\n", ebt)
	fmt.Printf("Profit After Tax : %.1f	\n", profit)
	fmt.Printf("Profit Ratio : %.1f%%\n", ratio)
}

const resultsFile = "financial_results.txt"

func writeResultsInFile(filename string, ebt, profit, ratio float64) error {
	// This is a placeholder for writing results to a file.
	resultText := fmt.Sprintf("Earnings Before Tax: %.2f\nProfit After Tax: %.2f\nProfit Ratio: %.2f%%\n", ebt, profit, ratio)
	err := os.WriteFile(filename, []byte(resultText), 0644)
	if err != nil {
		return err
	}
	return nil
}

func calculateFinancials(revenue, expenses, taxrate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt - (ebt * taxrate / 100)
	if profit != 0 {
		ratio = (ebt / profit) * 100
	}
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput < 0 {
		return 0, errors.New("input must be non-negative value")
	}
	return userInput, nil
}
