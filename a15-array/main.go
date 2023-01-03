package main

import "fmt"
import "strings"

func main() {
	
	// Basic Array
	var names[4] string

	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	for i := 0; i < len(names); i++ {
		fmt.Print(strings.ToUpper(names[i]), " ")
	}
	fmt.Println("\n")


	// Array initialization
	var fruits = [4]string{"apple", "orange", "banana", "melon"}

	for i := 0; i < len(fruits); i++ {
		fmt.Print(strings.ToUpper(fruits[i]), " ")
	}
	fmt.Println("\n")
	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element\t\t", fruits)
	fmt.Println("\n")

	// Vertical Array Declaration
	// var fruits2 [4]string

	// // cara horizontal
	// fruits2  = [4]string{"apple", "grape", "banana", "melon"}

	// // cara vertikal
	// fruits2  = [4]string{
	// 	"apple",
	// 	"grape",
	// 	"banana",
	// 	"melon",
	// }

	// Array Declaration without limit of array
	var numbers = [...]int{2, 3, 4, 5, 6}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	// Multidimensional Array
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	
	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

}