package main

import (
	"fmt"
	"strings"
)

func main() {

	// var avg = calculate(2, 3, 2, 9, 3, 5, 7, 2, 4, 8)

	// var msg = fmt.Sprintf("Hasil rata-rata: %.2f", avg)

	// fmt.Println(msg)

	var hobbies  = []string{"running", "reading", "programming"}
	yourHobbies("Hazlan Muhammad Qodri", hobbies...)
}

// 1 - Fungsi Variadic kayak slice
// func calculate(numbers ...int) float64 {

// 	var total int = 0

// 	for _, number := range numbers {
// 		total += number
// 	}

// 	var avg = float64(total) / float64(len(numbers))

// 	return avg

// }

func yourHobbies(name string, hobbies ... string) {

	var joinedHobbies = strings.Join(hobbies, " ")

	fmt.Printf("Hello, my name is: %s \n", name)
	fmt.Printf("My hobbies are: %s", joinedHobbies)
}