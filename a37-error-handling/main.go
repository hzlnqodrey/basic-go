package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Println("inputkan angka")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "  is a number")
	} else {
		fmt.Println(number, "  is not a number")
		fmt.Println(err.Error())
	}
}