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
	// for true {
	// 	fmt.Println("start!")
	// 	time.Sleep(1 * time.Second)
	// }
	
	// New Timer - Channeling and Blocking
	var timer = time.NewTimer(4 * time.Second)
	fmt.Println("start!")
	<-timer.C
	fmt.Println("finish")
	
	// time.AfterFunc - Channeling and Flex-Blocking
	var ch = make(chan bool)
	
	time.AfterFunc(4 * time.Second, func(){
		fmt.Println("expired")
		ch <- true
	})
	fmt.Println("start!")
	<- ch
	fmt.Println("finish")

	// time.After - Channel
	<- time.After(2 * time.Second)
	fmt.Println("finish")

	// Scheduler with TICKER
	done := make(chan bool)
	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(10 * time.Second) // wait 5 sec
		done <- true
	}()

	for {
		select {
			case <- done:
				ticker.Stop() // stopping the scheduler
				return
			case t := <- ticker.C:
				fmt.Println("Current time: ", t)
		}
	}
}