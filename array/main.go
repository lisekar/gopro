package main

import (
	"fmt"
)

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	var productNames [5]string = [5]string{"A Book", "A Carpet", "A Table", "A TV", "A speaker"}
	prices := [5]float64{10.0, 12.0, 30.0, 14.0, 50.0} // [int]datatype{item,item}
	fmt.Println(prices)
	productNames[1] = "A Table"
	fmt.Println(productNames[1])
	fmt.Println(productNames[3])

	featuredPrices := prices[1:3] // [1:3] - item 1 is included and item 3 is excluded
	fmt.Println(featuredPrices)

	slicePrices := prices[1:]
	fmt.Println("Array with removed item with postion given in bracket :: ", slicePrices)
	slicePrices[0] = 99
	fmt.Println(slicePrices)

	highlightedPrices := slicePrices[:1]
	fmt.Println(highlightedPrices)

	fmt.Println(prices, len(highlightedPrices), cap(highlightedPrices)) // modified data by slicePrices array, len() and cap()

	// add dynamic value in array - add and remove the elements
	dynamicPrice := []float64{100, 120, 130, 140, 150}
	fmt.Println("dynamic prices : ", dynamicPrice, "\n Selected element : ", dynamicPrice[4])
	fmt.Println(dynamicPrice[0:1])
	dynamicPrice[1] = 999
	// dynamicPrice[5] = 1000 // gives error - for creating a element in out of scope

	updatedPrices := append(dynamicPrice, 160) // create new array by appending in seperate value
	fmt.Println("updated prices : ", updatedPrices, dynamicPrice)
	dynamicPrice = append(dynamicPrice, 170)
	fmt.Println("updated prices : ", updatedPrices, dynamicPrice) // assign to parent array where you going to add
	dynamicPrice = dynamicPrice[1:]
	fmt.Println("dynamicPrice :::: removed [1] position element ", dynamicPrice)

	// get all the elements from the current list and separate element and added to existing element with '[list]...'
	discoutPrices := []float64{1000, 2000, 3000, 4000, 5000}
	dynamicPrice = append(dynamicPrice, discoutPrices...)
	fmt.Println("list add in another list : ", dynamicPrice)

}
