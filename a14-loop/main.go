package main

import "fmt"

func main() {
	// basic for
	for i := 0; i < 5; i++ {
		fmt.Printf("Angka 1st %d\n", i+1)
	}

	// for, only conditions without args
	var i = 0

	for i < 5 {
		fmt.Printf("Angka 2nd %d\n", i+1)
		i++
	} 

	// for, no args no conds
	var i2 = 0

	for {

		fmt.Printf("Angka 3rd %d\n", i2+1)

		i2++ 
		
		if i2 >= 5 {
			break
		}
		
	}
}