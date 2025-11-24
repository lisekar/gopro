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
	// expectedReturnRate := 5.5. // using short declaration , shorter notation only assinged inside the main function
	// var years float64 = 10  // using var keyword

	// fmt.Scanf("%f %f %f", &investmentAmount, &expectedReturnRate, &years)
	outputText("Investment Calculator Application \n")

	fmt.Print("Enter your Investment amount : ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter your Expected Return Rate : ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the number of Years : ")
	fmt.Scan(&years)

	result := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	fmt.Printf("Calculated Future Value result : %.2f\n", result)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureInflationValue := futureValue / math.Pow(1+inflationRate/100, years) // using const keyword
	// fmt.Println("Future Value of Investment : ", futureValue)
	// fmt.Println("Future Value Adjusted for Inflation : ", futureInflationValue)
	// formattedFV := fmt.Sprintf("Future Value of Investment : %.2f\n", futureValue)                    // formatting output
	// formattedFIV := fmt.Sprintf("Future Value Adjusted for Inflation : %.2f\n", futureInflationValue) // formatting output
	// Alternatively, you could also use:
	// fmt.Printf("Future Value of Investment : %v\nFuture Value Adjusted for Inflation : %v\n", futureValue, futureInflationValue)
	// fmt.Printf("Future Value of Investment : %v\nFuture Value Adjusted for Inflation : %v\n", futureValue, futureInflationValue)
	fmt.Printf(`Future Value of Investment : %.2f\n Future Value Adjusted for Inflation : %.2f`, futureValue, futureInflationValue) // using a single fmt.Printf statement with line break using backticks
	// fmt.Print(formattedFV, formattedFIV) // using a single fmt.Print statement

}

func outputText(text string) {
	// This is a placeholder for output text function.
	fmt.Print(text)
}

func calculateFutureValue(principal, rate, time float64) float64 {
	// This is a placeholder for future value calculation function.
	fv := principal * math.Pow(1+rate/100, time)
	return fv
}
