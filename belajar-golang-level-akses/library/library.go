package library

import "fmt"

var Student = struct {
	Name  string
	Grade int
}{}

func init() {
	Student.Name = "Hazlan Qodrey"
	Student.Grade = 100

	fmt.Println("----> library/library.go imported")
}