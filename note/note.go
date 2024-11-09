package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	Title   	string
	Content 	string
	CreatedAt   time.Time
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
