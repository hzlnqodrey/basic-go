package main

import (
	. "belajar-golang-level-akses/library"
	f "fmt"
)

func main() {
	
	var s1 = Student{"Hazlan", 90}
	
	f.Println("Name: ", s1.Name)
	f.Println("Grade: ", s1.Grade)

}