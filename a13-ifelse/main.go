package main

import "fmt"

func main() {
	// if else
	var point = -1

	if point == 8 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus, nilai anda: %d\n", point)
	}

	// variable temporary
	var point2 = 8900.0

	if percent := point2 / 100; percent >= 100 {
		fmt.Printf("%.1f%s Perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s Good!\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s Not Bad!\n", percent, "%")
	}

	// switch case
	var point3 uint = 6

	switch point3 {
	case 8:
		fmt.Printf("perfect")
	case 7:
		fmt.Printf("awesome")
	case 6:
		fmt.Printf("not bad")
	default:
		fmt.Printf("gg")
	}
}