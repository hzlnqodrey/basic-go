package main

import (
	"fmt"
)

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

func main() {
	var data = student{
		name:        "wick",
		height:      182.5,
		age:         26,
		isGraduated: false,
		hobbies:     []string{"eating", "sleeping"},
	}

	// %b - STRING NUMERIK BINER
	fmt.Printf("%b\n", data.age)

	// %c - STRING UNICODE
	fmt.Printf("%c\n", 1400)
	fmt.Printf("%c\n", 1235)
	
	// %d - STRING NUMERIK BASIS 10
	fmt.Printf("%d\n", 1235)
	fmt.Printf("%d\n", data.age)
	
}
