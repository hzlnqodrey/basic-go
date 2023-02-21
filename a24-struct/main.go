package main

import (
	"fmt"
)

type student struct {
	name string
	grade int
}

func main() {

	// var s1 student

	// s1.name = "Hazlan Muhammad Qodri"
	// s1.grade = 90

	// var s2 = student{"ethan", 20}

	// var s3 = student{name: "jason"}

	// fmt.Println("name : ", s1.name)
	// fmt.Println("grade : ", s1.grade)

	// fmt.Println("name : ", s2.name)
	// fmt.Println("name : ", s3.name)

	// Variable Struct Pointer
	var s1 = student{
		name: "hazlan",
		grade: 80,
	}

	var s2 *student = &s1
	fmt.Println("student 1, name: ", s1.name)
	fmt.Println("student 2, name: ", s2.name)

	s2.name = "rivano"
	fmt.Println("student 1, name: ", s1.name)
	fmt.Println("student 2, name: ", s2.name)
}