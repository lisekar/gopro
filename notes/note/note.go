package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// create struct type
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note *Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n%v", note.Title, note.Content)
}

func (note *Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	// convert to json from text file
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

// create constructor for Note struct
func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("text and content should not be empty")
	}
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
