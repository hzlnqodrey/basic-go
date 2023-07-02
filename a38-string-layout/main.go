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

	// %e %E - SCIENTIFIC NOTABLE
	fmt.Printf("%e\n", data.height)

	// %F - DECIMAL
	fmt.Printf("%f\n", data.height)
	// 182.500000
	fmt.Printf("%.9f\n", data.height)
	// 182.500000000
	fmt.Printf("%.2f\n", data.height)
	// 182.50
	fmt.Printf("%.f\n", data.height)
	// 182

	// %G - FREE FLOW DECIMAL
	fmt.Printf("%g\n", 0.123123123123)
	// 0.123123123123

	// %O - OCTAL BASE 8
	fmt.Printf("%o\n", data.age)
	// 32

	// %p - FORMAT POINTER
	fmt.Printf("%p\n", &data.name)

	fmt.Printf("%q\n", `" name \ height "`)
	// "\" name \\ height \""
}
