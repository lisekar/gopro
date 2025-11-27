package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (note *Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n%v", note.title, note.content)
}

func (note *Note) Save() error {
	fileName := strings.ReplaceAll(note.title, "", "_")
	fileName = strings.ToLower(fileName)
	// convert to json from text file
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("text and content should not be empty")
	}
	return &Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
