package main

import (
	"fmt"
	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	n1, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n1)

}

func getNoteData() (string, string) {
	title := getUserData("Title: ")
	content := getUserData("Content: ")

	return title, content

}

func getUserData(prompt string) string {
	fmt.Printf("%v: ", prompt)
	var value string
	fmt.Scanln(&value)
	return value
}