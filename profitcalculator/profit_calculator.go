package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxrate float64

	fmt.Print("Enter your Profit amount : ")
	fmt.Scan(&revenue)

	fmt.Print("Enter your Profit Rate : ")
	fmt.Scan(&expenses)

	fmt.Print("Enter the taxrate : ")
	fmt.Scan(&taxrate)

	var ebt = revenue - expenses
	profit := float64(ebt) - (1 * taxrate / 100)
	ratio := (ebt / profit) * 100

	fmt.Println("Earnings Before Tax : ", ebt)
	fmt.Println("Profit After Tax : ", profit)
	fmt.Println("Profit Ratio : ", ratio)
}
