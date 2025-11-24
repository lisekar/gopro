package main

import (
	"fmt"
	"math"
)

func main() {
	// This is a placeholder for the investment calculator application.
	const inflationRate = 2.5                                        // using const keyword
	investmentAmount, expectedReturnRate, years := 1000.0, 5.5, 10.0 // use short declaration
	// Alternatively, you could declare them like this:
	// var investmentAmount float64 = 1000.0 // using var keyword
	// expectedReturnRate := 5.5. // using short declaration
	// var years float64 = 10  // using var keyword

	// fmt.Scanf("%f %f %f", &investmentAmount, &expectedReturnRate, &years)

	fmt.Print("Enter your Investment amount : ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter your Expected Return Rate : ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the number of Years : ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureInflationValue := futureValue / math.Pow(1+inflationRate/100, years) // using const keyword
	fmt.Println("Future Value of Investment : ", futureValue)
	fmt.Println("Future Value Adjusted for Inflation : ", futureInflationValue)
}
