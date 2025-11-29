package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// create struct type
type Todo struct {
	Text string `json:"title"`
}

func (todo *Todo) Display() {
	fmt.Printf("Your note titled %v : \n", todo.Text)
}

func (todo *Todo) Save() error {
	// fileName := strings.ReplaceAll(todo.Title, " ", "_")
	// fileName = strings.ToLower(fileName) + ".json"
	fileName := "todo.json"
	// convert to json from text file
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

// create constructor for Note struct
func New(content string) (*Todo, error) {
	if content == "" {
		return &Todo{}, errors.New("text should not be empty")
	}
	return &Todo{
		Text: content,
	}, nil
}
