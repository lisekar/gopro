package main

import (
	"fmt"
)

type funcTransform func(int) int // make customdatatype - with alias
// type sampleFunction func(int, []string, map[string][]int) ([]int, string) // create alias as per our logic

func main() {

	numbers := []int{1, 2, 3, 4, 5}
	tripNumbers := []int{5, 1, 3, 2, 6}
	doubNumbers := []int{1, 4, 5, 7, 8, 9}
	doubled := transformNumbers(&numbers, double) // parameter with function
	triple := transformNumbers(&numbers, triple)  // parameter with function
	fmt.Println("Doubled Number : ", doubled)
	fmt.Println("Tripled Number : ", triple)

	tranformNum := getTransformNumber(&tripNumbers)
	showTransformNumbers := transformNumbers(&tripNumbers, tranformNum)
	fmt.Println("showTransformNumbers :: ", showTransformNumbers)

	tranformDoubleNum := getTransformNumber(&doubNumbers)
	showTransformNumbersDouble := transformNumbers(&doubNumbers, tranformDoubleNum)
	fmt.Println("showTransformNumbers Double :: ", showTransformNumbersDouble)

	// Anonymous function - doesn't have any function name
	anonymousFunc := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
	fmt.Println("Anonymous Function :: ", anonymousFunc)

	// calling closure method
	numNoubled := createClosureTranform(2) // calling closure method
	numTripled := createClosureTranform(3) // calling a closure method

	double := transformNumbers(&numbers, numNoubled) // passing a function as param
	fmt.Println("double Function :: ", double)
	trilpe := transformNumbers(&numbers, numTripled) // passing a function as param
	fmt.Println("trilpe Function :: ", trilpe)

	// calling sample closure
	getClosureCal := sampleClosure(5)(2)
	fmt.Println("Simple clouser :: ", getClosureCal)

	// Anonymous closure method
	fmt.Println("Closure using anonymous function : ", anonymClosure)

}

// function that can receive parameter as another function
func transformNumbers(numbers *[]int, transform funcTransform) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		// fmt.Println(" Val : ", val)
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

// function that return as another function - reduce quite number of lines in code
func getTransformNumber(numbers *[]int) funcTransform {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

// Closure
func createClosureTranform(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

// simple closure function
func sampleClosure(a int) func(int) int {
	return func(b int) int {
		return a * b
	}
}

// anonymous closure function
var anonymClosure = func(a int) func(int) int {
	return func(b int) int {
		return a * b
	}
}(5)(2)
