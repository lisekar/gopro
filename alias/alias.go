package main

import "fmt"

// create a custom type
type livistr string

func (text livistr) log() {
	fmt.Println(text)
}
func main() {
	var name livistr = "Hi, Livi"
	name.log()
}
