package main

import (
	"errors"
	"fmt"
	"strings"
)

// 2 - Custom Error
func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}

	return true, nil
}

func main() {
	var name string
	fmt.Println("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Hai", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}
