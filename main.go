package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"example.com/note/note"
	"example.com/note/todo"
)
type displayer interface {
	Display()
}

type outputtable interface {
	displayer
	Save() error
}

func main() {
	title, content := getNoteData()
	note1, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return 
	}

	text := getUserData("Please enter text for you todo: ")
	fmt.Println()
	todo1, err := todo.New(text)
	if err != nil {
		fmt.Println(err)
		return 
	}

	err = outputData(todo1)

	if err != nil {
		fmt.Println(err)
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(note1)

}

func getNoteData() (string, string) {
	title := getUserData("Title: ")
	content := getUserData("Content: ")
	fmt.Println()
	return title, content

}

func getUserData(prompt string) string {
	fmt.Printf("%v: ", prompt)
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	return strings.TrimSpace(value)
}

func outputData(data outputtable) error {
	data.Display()
	err := data.Save()
	return err
}

