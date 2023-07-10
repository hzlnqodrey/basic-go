package main

import (
	"fmt"
	"math/rand"
	"time"
)

// String Random
var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {

	// Unique Random Seed
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())

	a := randomString(10)

	fmt.Println(a)
}

func randomString(length int) string {
	b := make([]rune, length)

	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}