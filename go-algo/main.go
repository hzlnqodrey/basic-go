package main

import (
	"fmt"
	"log"
)

func bubbleSort(arr []int) []int {

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

	return arr
}

func binary_search(arr []int, value int) (bool, int) {

	lower := 0
	upper := len(arr) - 1
	var middle int
	for lower <= upper {

		middle = (lower + upper) / 2

		if value > arr[middle] {
			lower = middle + 1
		} else if value < arr[middle] {
			upper = middle - 1
		} else if value == arr[middle] {
			return true, middle
		}

	}

	return false, middle
}

func main() {

	// random list
	var nums = []int{8, 4, 9, 3, 21, 11}
	fmt.Println(nums)

	// bubble sort
	sortedArr := bubbleSort(nums)

	fmt.Println(sortedArr)

	// binary search
	var i int

	if _, err := fmt.Scanf("%d", &i); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("i: %d\n", i)

	value := i

	ketemu, index := binary_search(sortedArr, value)

	if ketemu != false {
		fmt.Printf("Hasil ketemu pada index %d\n", index)
	} else {
		fmt.Printf("Angka tidak ditemukan pada list %v\n", sortedArr)
	}
}
