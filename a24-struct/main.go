package main

import (
	"fmt"
)

type student struct {
	name string
	grade int
}

func main() {

	var s1 student

	s1.name = "Hazlan Muhammad Qodri"
	s1.grade = 90

	fmt.Println("name : ", s1.name)
	fmt.Println("grade : ", s1.grade)
}