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

	// %Q - ESCAPE FORMAT
	fmt.Printf("%q\n", `" name \ height "`)
	// "\" name \\ height \""

	// %S - NORMAL STRING FORMAT
	fmt.Printf("%s\n", data.hobbies)

	// %t - BOOL FORMAT
	fmt.Printf("%t\n", data.isGraduated)
	// false

	// %T - TIPE DATA FORMAT
	fmt.Printf("%T\n", data.name)
	// string
	fmt.Printf("%T\n", data.height)
	// float64
	fmt.Printf("%T\n", data.age)
	// int32
	fmt.Printf("%T\n", data.isGraduated)
	// bool
	fmt.Printf("%T\n", data.hobbies)
	// []string

	// %V - ANY FORMAT [FREE]
	fmt.Printf("%v\n", data)
	// {wick 182.5 26 false [eating sleeping]}

	// %+v - VARIABLE + VALUE
	fmt.Printf("%+v\n", data)
	// {name:wick height:182.5 age:26 isGraduated:false hobbies:[eating sleeping]}

	// %#v - STRUCT FORMAT
	fmt.Printf("%#v\n", data)
	// main.student{name:"wick", height:182.5, age:26, isGraduated:false, hobbies:

	// %x- HEXA [16]
	fmt.Printf("%x\n", data.age)
	// 1a
	var d = data.name
	fmt.Printf("%x%x%x%x\n", d[0], d[1], d[2], d[3])
	// 7769636b
	fmt.Printf("%x\n", d)
	// 7769636b

	// %% [yeah]
	fmt.Printf("%%\n")
	// %

}
