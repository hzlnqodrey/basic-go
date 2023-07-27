package main

import (
	"fmt"
)

func main() {
	// deklarasi pointer golang
	// var number *int
	// var name *string

	// penggunaan pointer
	// var numberA int = 4
	// var numberB *int = &numberA

	// fmt.Println("numberA (value): ", numberA)
	// fmt.Println("numberA (address): ", &numberA)

	// fmt.Println("numberB (value): ", *numberB)
	// fmt.Println("numberB (address): ", numberB)

	// numberA = 5

	// fmt.Println("numberA (value): ", numberA)
	// fmt.Println("numberA (address): ", &numberA)

	// fmt.Println("numberB (value): ", *numberB)
	// fmt.Println("numberB (address): ", numberB)


	// Pointer as Parameter
	var number int = 4

	fmt.Println("Before: ", number)

	change(&number, 10)

	fmt.Println("After: ", number)
}

func change(original *int, value int) {
	*original = value
}