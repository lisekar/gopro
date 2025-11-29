package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {

	// create a list
	// cousreRating := map[string]float64{}			// can create and initialize map
	cousreRating := make(map[string]float64, 3) // we can also assign and initialize with make method

	// reallocate memory without initialize cap lentgh (maximum no. in this array)
	cousreRating["go"] = 4.7
	cousreRating["node"] = 4.5
	cousreRating["mongo"] = 4.4
	cousreRating["React"] = 4.9

	fmt.Println(cousreRating)

	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)

	// update a existing Key Value pair - mutate key value pair
	websites["Google"] = "https://www.google.com"
	websites["LinkedIn"] = "https://Linkedin.com"
	fmt.Println(websites)

	// Delete items from list
	delete(websites, "Google")
	fmt.Println(websites)
}
