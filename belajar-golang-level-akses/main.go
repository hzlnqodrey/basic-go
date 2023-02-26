package main

import (
	"belajar-golang-level-akses/library"
	"fmt"
)

func main() {
	
	var s1 = library.Student{"Hazlan", 90}
	
	fmt.Println("Name: ", s1.Name)
	fmt.Println("Grade: ", s1.Grade)

}