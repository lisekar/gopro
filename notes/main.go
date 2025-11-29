package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/livisekar/gopro/notes/note"
	"github.com/livisekar/gopro/notes/todo"
)

type saver interface {
	Save() error
}

type output interface {
	saver
	Display()
}

// access any type value in interface {} method
func add[T int | float64 | string](a, b T) T { // generic method - with placeholder [T] - it should be anytype
	// without Generic method.
	// aInt, aIsInt := a.(int)
	// bInt, bIsInt := b.(int)
	// if aIsInt && bIsInt {
	// 	return aInt + bInt
	// }
	// return nil
	return a + b
}

func main() {

	// call the generic method
	result := add(1, 2)
	fmt.Println(result)

	printSomething(23)      // integer
	printSomething(1.5)     // float32
	printSomething("Hello") // string

	title, content, err := getNoteData()
	if err != nil {
		fmt.Print(err)
		return
	}

	var userNote *note.Note
	userNote, err = note.New(title, content)
	if err != nil {
		fmt.Println("invalid input")
	}
	outputData(userNote)
	// userNote.Display()
	// userNote.Save()

	// var todoText *todo.Todo
	todoText, err := getUserInput("Todo text :")
	fmt.Println(todoText)
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Print(err)
		return
	}

	// todo.Display()
	// todo.Save()

	// display data by embeded interface
	outputData(todo)
	// if err != nil {
	// 	fmt.Println("Saving note failed")
	// 	return
	// }
	err = saveData(todo) // call this function by interface
	if err != nil {
		return
	}

	err = saveData(userNote) // call this function by interface
	if err != nil {
		return
	}

	// fmt.Println("Note saved !")

}

func printSomething(value interface{}) {
	// Alternate syntax
	typeVal, checkType := value.(int)
	if checkType {
		fmt.Println("Integer : ", typeVal)
	}

	floatVal, checkType := value.(float64)
	if checkType {
		fmt.Println("Integer : ", floatVal)
	}

	stringVal, checkType := value.(string)
	if checkType {
		fmt.Println("Integer : ", stringVal)
	}

	// switch value.(type) {        // syntax for switch statement
	// case int:
	// 	fmt.Println("Integer : ", value)
	// case float64:
	// 	fmt.Println("float64 : ", value)
	// case string:
	// 	fmt.Println("string : ", value)
	// }
}

// by accesing with embeded interface
func outputData(data output) error {
	data.Display()
	return saveData(data)
}

// method accesing with intrface
func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}
	fmt.Println("Note saved !")
	return nil
}

// var todoText *todo.Todo

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note title : ")
	if err != nil {
		// fmt.Print(err)
		return "", "", err
	}

	content, err := getUserInput("Note content : ")
	if err != nil {
		// fmt.Print(err)
		return "", "", err
	}

	return title, content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Printf("%v ", prompt)
	/*
		// only can get a short or single value by using scanln

			// var value string
			// fmt.Scanln(&value)
			// if value == "" {
			// 	return "", errors.New("invalid input")
			// }
			// return value, nil
	*/
	reader := bufio.NewReader(os.Stdin) // reade user input text from command line
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("invalid input")
	}
	text = strings.TrimSuffix(text, "\n") // remove new line by this method using strings package
	text = strings.TrimSuffix(text, "\r") // remove new line by this method using strings package
	return text, nil
}
