package main

import(
	"fmt"
	"time"
)

func main(){
	var time1 = time.Now()
	fmt.Printf("time %v\n", time1)
	
	var time2 = time.Date(2023, 7, 12, 11, 27, 10, 0, time.UTC)
	fmt.Printf("time %v\n", time2)

}