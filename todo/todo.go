package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string
}

func New(data string) (Todo, error) {

	if data == "" {
		return Todo{}, errors.New("unable to get the text for todo list")
	}

	return Todo {
		Text: data,
	}, nil
}


func (todo Todo) Display() {
	fmt.Printf("Here is your Todo:\n")
	fmt.Printf("Text: %v\n\n", todo.Text,)
}


func (t Todo) Save() error {
	jsonData, err := json.MarshalIndent(t, "", " ")

	if err != nil {
		return fmt.Errorf("could not marshal note to JSON: %w", err)
	}

	err = os.WriteFile("todo.json", jsonData, 0644)

	if err != nil {
		return fmt.Errorf("could not write to file: %w", err)
	}

	fmt.Printf("Successfully saved Todo to %s\n", "todo.json")
	return nil
}
