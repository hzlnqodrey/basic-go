package main

import (
	"fmt"
	"runtime"
)

// 2 - Channel sebagai tipe data parameter
func printMessages(what chan string) {
	fmt.Println(<-what)
}

func main() {
	// 1 - penerapan channel
	// runtime.GOMAXPROCS(2)

	// var message = make(chan string)
	
	// var sayHelloTo = func(who string) {
	// 	var data = fmt.Sprintf("Hello %s", who)
	// 	message <- data
	// }

	// go sayHelloTo("hazlan qodri")
	// go sayHelloTo("rivano atk")
	// go sayHelloTo("rifki afattah")

	// var message1 = <- message
	// fmt.Println(message1)

	// var message2 = <- message
	// fmt.Println(message2)
	
	// var message3 = <- message
	// fmt.Println(message3)

	// 2 - Channel sebagai tipe data parameter
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	for _, each := range []string{"Hazlan", "Rifki", "Rivano"} {
		go func(who string) {
			var data = fmt.Sprintf("Hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessages(messages)
	}
}
