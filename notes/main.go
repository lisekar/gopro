package main

import (
	"bufio"
	"fmt"
	"notes/note"
	"os"
	"strings"
)

func main() {
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
	userNote.Display()
}

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
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("invalid input")
	}
	text = strings.TrimSuffix(text, "\n") // remove new line by this method using strings package
	text = strings.TrimSuffix(text, "\r") // remove new line by this method using strings package
	return text, nil
}
