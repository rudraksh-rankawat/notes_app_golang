package note

import (
	"errors"
	"fmt"
)

type Note struct {
	title   string
	content string
}

func New(title, content string) (Note, error) {

	err := errors.New("unable to get title or content")

	if title == "" || content == "" {
		fmt.Println(err)
		return Note{}, err
	}

	return Note{
		title: title,
		content: content,
	}, nil
}
