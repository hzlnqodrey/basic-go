package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64
	diameter = 15
	
	var area, circumferences = calculate(diameter)

    fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
    fmt.Printf("keliling lingkaran\t: %.2f \n", circumferences)

}

// predefined 2 return

func calculate(d float64) (area, circumference float64) {
	area = math.Pi * math.Pow(d / 2, 2)
	circumference = math.Pi * d

	return
}