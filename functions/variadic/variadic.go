package main

import "fmt"

func main() {

	numbers := []int{1, 10, 15}

	showSum := sumup(numbers)
	fmt.Println("Show sum : ", showSum)

	// by using variadic method
	showSumVariadic := sumupVariadic(1, 2, 3, 4, 5, 6, 7, 8, -9) // use our own list params
	fmt.Println("Show sum : ", showSumVariadic)

	// by using variadic method - with custom start value of ur own
	showSumVariadicSplit := sumupVariadicSplit(1, 2, 3, 4, 5) // use our own list params
	fmt.Println("Show sum : ", showSumVariadicSplit)

	// spliting up to slice
	showSumVariadicSplitSlice := sumupVariadicSplit(1, numbers...) // slice params
	fmt.Println("Show sum : ", showSumVariadicSplitSlice)

}

func sumup(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val // sum = sum+1
	}
	return sum
}

func sumupVariadic(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val // sum = sum+1
	}
	return sum
}

func sumupVariadicSplit(startVal int, numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val // sum = sum+1
	}
	return sum
}
