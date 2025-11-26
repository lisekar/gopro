package main

import (
	"fmt"
)

func main() {
	age := 32 // normal integer variable
	fmt.Println("Age:", age)

	var agePointer *int = &age // getting address of age variable

	fmt.Println("Age Pointer:", agePointer)              // printing pointer address
	fmt.Println("Age Value using Pointer:", *agePointer) // dereferencing pointer

	*agePointer = 35                                             // updating value using pointer
	fmt.Println("Updated Age Value using Pointer:", *agePointer) // dereferencing pointer after update
	fmt.Println("Age:", age)

	adultYears := getAdultYears(agePointer) // passing pointer to function
	fmt.Println("Years since Adult Age:", adultYears)

	setAgeToAdult(agePointer) // passing pointer to function
	printAge(agePointer)      // passing pointer to function

	updateAge(agePointer) // passing pointer to function
	fmt.Println("Updated Age:", age)
}

func getAdultYears(agePtr *int) int {
	if *agePtr >= 18 {
		return *agePtr - 18
	}
	return 0
}

func updateAge(agePtr *int) {
	*agePtr = *agePtr - 18
}

func setAgeToAdult(agePtr *int) {
	if *agePtr < 18 {
		*agePtr = 18
	}
}

func printAge(agePtr *int) {
	fmt.Println("Age is:", *agePtr)
}
