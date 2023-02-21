package main

import (
	"fmt"
	"strings"
)

// Alias fungsi
type FilterCallback func(string) bool

func main() {
	
	var data = []string{
		"wick",
		"jason",
		"ethan",
	}

	// filter data yg ada huruf o/O
	var containsO = filtered(data, func(s string) bool {
		return strings.Contains(s, "o")
	})

	// filter data yang berjumlah 5
	var contains5 = filtered(data, func(s string) bool {
		return len(s) == 5
	})

	// print data asli
	fmt.Printf("Data Asli \t\t: %v\n", data)

	// print data filter O
	fmt.Printf("Data Filter O \t\t: %v\n", containsO)

	// print data filter yang berjumlah 5 huruf
	fmt.Printf("Data Filter huruf \t\t: %v\n", contains5)
}

func filtered(data []string, callback FilterCallback) []string {
	var result []string

	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	return result
}