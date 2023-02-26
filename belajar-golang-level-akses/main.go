package main

import (
	. "belajar-golang-level-akses/library"
	f "fmt"
)

func main() {
	f.Printf("Name:  %s\n", Student.Name)
	f.Printf("Grade:  %d\n", Student.Grade)
}