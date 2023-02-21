package main

import "fmt"

func main() {

	var getMinMax = func(n []int) (int, int) {
		var min, max int

		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		
		return min, max
	}

	var numbers = []int{2, 9, 8, 10, 4, 2, 3, 7}

	var min, max = getMinMax(numbers)

	fmt.Printf("data: %v\nmin : %v\nmax: %v\n", numbers, min, max)
}