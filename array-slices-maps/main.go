package main

import "fmt"

type floatMap map[string]float64 // create an alias custom type

func (d floatMap) output() {
	fmt.Println(d)
}

func main() {

	// create array
	userNames := []string{} // only insert by append method
	// userNames[1] = "mann"  // not accepted in this method
	userNames = append(userNames, "Livi")
	userNames = append(userNames, "ven")

	fmt.Println(userNames)

	product := make([]string, 2, 5) // pre allocated memory and allocate new space.
	product[0] = "pro_1"
	product[1] = "pro_2"
	// product[3] = "pro_3"
	product = append(product, "Livi")
	product = append(product, "ven")

	fmt.Println(product)

	grade := make(floatMap, 3)
	grade["English"] = 4.5
	grade["Tamil"] = 4.9
	grade["Science"] = 4.7

	grade.output() // call alias custom type method

	// fmt.Println(grade)

	for index, user := range userNames {
		fmt.Println("index : ", index, "\n User : ", user)
	}

	for key, value := range grade {
		fmt.Println("index : ", key, "\n User : ", value)
	}

}
