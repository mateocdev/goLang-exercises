package main

import (
	"errors"
	"fmt"
)

func main() {
	title, body, err := getNoteData()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Title:", title)
	fmt.Println("Body:", body)
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note title:")

	if err != nil {
		fmt.Println("Error:", err)
		return "", "", err
	}

	body, err := getUserInput("Note body:")

	if err != nil {
		fmt.Println("Error:", err)
		return "", "", err
	}

	return title, body, nil
}

func getUserInput(promt string) (string, error) {
	fmt.Print(promt)
	var value string
	fmt.Scanln(&value)
	if value == "" {
		return "", errors.New("empty value")
	}
	return value, nil
}
