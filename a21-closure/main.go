package main

import "fmt"

func main() {

	// 1 - Closure
	// var getMinMax = func(n []int) (int, int) {
	// 	var min, max int

	// 	for i, e := range n {
	// 		switch {
	// 		case i == 0:
	// 			max, min = e, e
	// 		case e > max:
	// 			max = e
	// 		case e < min:
	// 			min = e
	// 		}
	// 	}

	// 	return min, max
	// }

	// var numbers = []int{2, 9, 8, 10, 4, 2, 3, 7}

	// var min, max = getMinMax(numbers)

	// fmt.Printf("data: %v\nmin : %v\nmax: %v\n", numbers, min, max)

	// 2 - Clousre IIFE
	
	var numbers  = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	// filtered number that below 3, delete them
	var filteredNumbers = func(min int) []int {

		var r []int

		for _, e := range numbers {
			if e < min {
				continue
			}
				
			r = append(r, e)
		}

		return r

	}(3)

	fmt.Printf("original number: %v\n", numbers)
	fmt.Printf("filtered number: %v\n", filteredNumbers)

}