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

	displayData(todo1)
	displayData(note1)
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

func displayData(data displayer) {
	data.Display()
}