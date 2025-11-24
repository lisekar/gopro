package main

import "fmt"

func main() {
	var profitAmount float64
	var profitRate float64
	var durationMonths float64

	fmt.Print("Enter your Profit amount : ")
	fmt.Scan(&profitAmount)

	fmt.Print("Enter your Profit Rate : ")
	fmt.Scan(&profitRate)

	fmt.Print("Enter the Duration in Months : ")
	fmt.Scan(&durationMonths)

	totalProfit := (profitAmount * profitRate * durationMonths) / (100 * 12)
	fmt.Println("Total Profit : ", totalProfit)
}
