package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	note "example.com/note/notes"
)

func main() {
	title, body := getNoteData()

	userNote, err := note.New(title, body)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Calling metod display from note.go
	userNote.Display()
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	body := getUserInput("Note body:")

	return title, body
}

func getUserInput(promt string) string {
	fmt.Print(promt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
