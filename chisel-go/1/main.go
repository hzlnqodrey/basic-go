package main

import (
	"fmt"
	"time"
)

func main() {
	var result int

	go func() {
		for i := 1; i <= 1000000; i++ {
			result += 1
		}
	}()

	go func() {
		for i := 1; i <= 1000000; i++ {
			result -= 1
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(result)
}