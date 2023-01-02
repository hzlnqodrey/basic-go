package main

import "fmt"

func main() {

	const firstName string = "john"
	fmt.Print("halo ", firstName, "!\n")

	// Inference constanta
	const lastName = "wick"
	fmt.Print("nice to meet you ", lastName, "!\n")

	// multi constanta
	const (
		square          = "kotak"
		isToday bool    = true
		numeric uint8   = 1
		floatNum        = 2.2
	)

	// this output will be the same
	const (
		a = "konstanta"
		b // konstanta
	)

	// determine ...
	const (
		today string = "senin"
		sekarang
		isToday2 bool = true
	)

	// multi variable in one line
	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"
}
