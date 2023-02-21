package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	// var names = []string{"John", "Wick"}
	// printMessage("halo", names)

	rand.Seed(time.Now().Unix())
	var RandomNumber int

	RandomNumber = randomWithRange(2, 10)
	fmt.Println("random number: ", RandomNumber)

	RandomNumber = randomWithRange(2, 10)
	fmt.Println("random number: ", RandomNumber)

	RandomNumber = randomWithRange(2, 10)
	fmt.Println("random number: ", RandomNumber)

}

// // 1 - void func
// func printMessage(message string, arr_names []string) {
// 	var Joined = strings.Join(arr_names, " ")
// 	fmt.Println(message, Joined)
// }

// 2 - return func
func randomWithRange(min, max int) int {
	var value = rand.Int() % (max - min + 1) + min
	return value
}

// // Deklarasi Parameter bertipe data sama
// func nameOfFunc(paramA type, paramB type, paramC type) returnType
// func nameOfFunc(paramA, paramB, paramC type) returnType

// func randomWithRange(min int, max int) int
// func randomWithRange(min, max int) int
