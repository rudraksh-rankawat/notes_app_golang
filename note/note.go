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
	Title   	string      `json:"title"`
	Content 	string		`json:"content"`
	CreatedAt   time.Time	`json:"created_at"`
}

func New(title, content string) (Note, error) {

	err := errors.New("unable to get title or content")

	if title == "" || content == "" {
		fmt.Println(err)
		return Note{}, err
	}

	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) Display() {
	fmt.Printf("Here is your Note:\n")
	fmt.Printf("Title: %v\nContent: %v\n", n.Title, n.Content)
}


func (n Note) Save() error {
	fileName := strings.ToLower(strings.ReplaceAll(n.Title, " ", "_")) + ".json"
	jsonData, err := json.MarshalIndent(n, "", " ")

	if err != nil {
		return fmt.Errorf("could not marshal note to JSON: %w", err)
	}

	err = os.WriteFile(fileName, jsonData, 0644)

	if err != nil {
		return fmt.Errorf("could not write to file: %w", err)
	}

	fmt.Printf("Successfully saved Note to %s\n", fileName)
	return nil
}