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

// 3 - Recover
func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error Occured", r)
	} else {
		fmt.Println("Application Running Flawlessly")
	}
}

func main() {

	data := []string{"superman", "aquaman", "wonder woman"}

	for _, each := range data {

		func() {

			// recover untuk IIFE dalam perulangan
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic occured on looping", each, "| message:", r)
				} else {
					fmt.Println("Application running perfectly")
				}
			}()

			panic("some error happen")
		}()

	}

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
