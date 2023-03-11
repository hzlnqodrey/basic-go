package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 1 - penerapan channel
	runtime.GOMAXPROCS(2)

	var message = make(chan string)
	
	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("Hello %s", who)
		message <- data
	}

	go sayHelloTo("hazlan m qodri")
	go sayHelloTo("rivano atk")
	go sayHelloTo("rifki afattah")

	var message1 = <- message
	fmt.Println(message1)

	var message2 = <- message
	fmt.Println(message2)
	
	var message3 = <- message
	fmt.Println(message3)
}
