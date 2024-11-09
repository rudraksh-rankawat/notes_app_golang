package todo

import (
	"errors"
	"fmt"
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
