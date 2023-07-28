package main

import (
	"fmt"
	"time"
)

func main(){

	// time.sleep
	fmt.Println("Hello!")
	time.Sleep(4 * time.Second)
	fmt.Println("After 4 seconds!")

	// Scheduler
	for true {
		fmt.Println("start!")
		time.Sleep(1 * time.Second)
	}
}