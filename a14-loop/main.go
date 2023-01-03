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

	// FOR | break -  and continue
	for number := 0; number < 10; number++ {

		fmt.Printf("Number: %d\n", number)

		if number % 2 == 1 {
			fmt.Printf("Number is odd, continue \n")
			continue
		}

		if number > 5 {
			fmt.Printf("Number berada di iterasi ke 5, break \n")
			break
		}

	}

	// Nested Loop
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
	
		fmt.Println()
	}

	// Label in Loop
	outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
