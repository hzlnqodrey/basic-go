package main

import (
	"fmt"
)

func main() {

	var secret interface{}

	secret = "Hazlan Muhammad Qodri"
	fmt.Println(secret)

	secret = []string{"Banana", "Melon"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)
}