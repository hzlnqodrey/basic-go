package main

import "fmt"
// import "strings"

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0])


	// Perbedaan Slice dan Array

	var fruitsA = []string{"banana", "melon"}
	var fruitsB = [2]string{"strowberry", "watermelon"}
	var fruitsC = [...]string{"bluberry", "pomegranate"}

	for i, fruits := range fruitsA {
		fmt.Printf("Fruits %d: %s", i+1, fruits)
		fmt.Println(" ")
	}

	for i, fruits := range fruitsB {
		fmt.Printf("Fruits %d: %s", i+1, fruits)
		fmt.Println(" ")
	}

	for i, fruits := range fruitsC {
		fmt.Printf("Fruits %d: %s", i+1, fruits)
		fmt.Println(" ")
	}

	// akses slice
	var newFruits = fruitsA[0:2]
	fmt.Println(newFruits)

	
	var fruitsSS = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruitsSS[0:3]
	var bFruits = fruitsSS[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	var fruitsFUNC = []string{"pisang", "apel", "jeruk", "melon"}
	fmt.Println(len(fruitsFUNC))
	fmt.Println(cap(fruitsFUNC))

	// 

	fmt.Println(" ")

	var fruitsZ = []string{"pisang", "apel", "jeruk"}

	var aFruitsZ = fruitsZ[0:2]

	fmt.Println(cap(aFruitsZ))
	fmt.Println(len(aFruitsZ))

	fmt.Println(fruitsZ)
	fmt.Println(aFruitsZ)

	var cFruitsZ = append(aFruitsZ, "papaya")

	fmt.Println(fruitsZ)
	fmt.Println(aFruitsZ)
	fmt.Println(cFruitsZ)


	// Copy Function
	dst := make([]string, 3)
	src := []string{"watermelon", "papaya", "banana", "sugar"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)

	// Pengaksesan 3 Index Slicing 
	var fruits = []string{"apple", "grape", "banana"}
	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2] // untuk menentukan kapasitas

	fmt.Println(fruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits)) // len: 3
	fmt.Println(cap(fruits)) // cap: 3

	fmt.Println(aFruits)      // ["apple", "grape"]
	fmt.Println(len(aFruits)) // len: 2
	fmt.Println(cap(aFruits)) // cap: 3

	fmt.Println(bFruits)      // ["apple", "grape"]
	fmt.Println(len(bFruits)) // len: 2
	{fmt.Println(cap(bFruits)) // cap: 2}

}