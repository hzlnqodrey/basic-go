package main

import (
	"fmt"
)

func bubbleSort(arr []int) {

	n := len(arr)
	for i := 0; i < n; i++ {
		var swapped bool = false

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if swapped != true {
			break
		}

	}
}

func main() {

	// random list
	var nums = []int{8, 4, 9, 3, 21, 11}
	fmt.Println(nums)
	
	// bubble sort
	bubbleSort(nums)

	fmt.Println(nums)

	// binary search
	
}
