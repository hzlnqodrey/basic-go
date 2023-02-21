package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}
// type student struct {
// 	grade int
// 	person
// }

func main() {

	// 1 - Struct Declaration
	// var s1 student

	// s1.name = "Hazlan Muhammad Qodri"
	// s1.grade = 90

	// var s2 = student{"ethan", 20}

	// var s3 = student{name: "jason"}

	// fmt.Println("name : ", s1.name)
	// fmt.Println("grade : ", s1.grade)

	// fmt.Println("name : ", s2.name)
	// fmt.Println("name : ", s3.name)

	// 2 - Variable Struct Pointer
	// var s1 = student{
	// 	name: "hazlan",
	// 	grade: 80,
	// }

	// var s2 *student = &s1
	// fmt.Println("student 1, name: ", s1.name)
	// fmt.Println("student 2, name: ", s2.name)

	// s2.name = "rivano"
	// fmt.Println("student 1, name: ", s1.name)
	// fmt.Println("student 2, name: ", s2.name)

	// 3 - Embedded Struct
	// var s1 = student{}
	// s1.name = "Hazlan"
	// s1.age = 22
	// s1.grade = 80

	// fmt.Println("Name: ", s1.name)
	// fmt.Println("Age: ", s1.age)
	// fmt.Println("Age: ", s1.person.age)
	// fmt.Println("Grade: ", s1.grade)

	// 4 - Pengisian Sub-struct
	// var p1 = person{
	// 	name: "Hazlan",
	// 	age: 22,
	// }

	// var s1 = student{
	// 	person: p1,
	// 	grade: 85,
	// }

	// fmt.Println("Name: ", s1.name)
	// fmt.Println("Age: ", s1.age)
	// fmt.Println("Age: ", s1.person.age)
	// fmt.Println("Grade: ", s1.grade)

	// 5 - Anonymous Struct
	var s1 = struct{
		person
		grade int
	}{}

	s1.person = person{"Hazlan", 23}
	s1.grade = 100

	fmt.Println("name  :", s1.person.name)
    fmt.Println("age   :", s1.person.age)
    fmt.Println("grade :", s1.grade)

	fmt.Println("Name: ", s1.name)
	fmt.Println("Age: ", s1.age)
	fmt.Println("Age: ", s1.person.age)
	fmt.Println("Grade: ", s1.grade)
}