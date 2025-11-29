package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {

	// task 1 -  creating a new array
	hobbies := []string{"Sports", "Cooking", "Reading"}
	fmt.Println(hobbies)

	// task 2 - display element
	fmt.Println(hobbies[0])   // display the first element
	fmt.Println(hobbies[1:3]) // disply the second and third element

	// task 3
	mainhobbies := hobbies[:2]
	fmt.Println(mainhobbies)

	// task 4
	fmt.Println(cap(mainhobbies))
	mainhobbies = mainhobbies[1:3]
	fmt.Println(mainhobbies)

	// task 5 - slice array
	courseGoals := []string{"Learn Go !", "Learn all the basics"}
	fmt.Println("course Goals : ", courseGoals)
	// courseGoals[0] = "Learn Go !"
	courseGoals[1] = "Learn all the details"
	fmt.Println(courseGoals)

	// task 6
	courseGoals = append(courseGoals, "Learn all the basics")
	fmt.Println(courseGoals)

	// task 7
	products := []Product{
		{"first-product", "A first product", 199},
		{"second-product", "A second product", 280},
	}
	fmt.Println(products)

	newProduct := Product{ // crated new product
		"third-product",
		"A third product",
		120,
	}
	products = append(products, newProduct) // append in exisitng struct
	fmt.Println("Products are : ", products)

}
