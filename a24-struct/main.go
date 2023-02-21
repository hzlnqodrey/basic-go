package main

import (
	"fmt"
)

// type person struct {
// 	name string
// 	age int
// }

// type student struct {
// 	grade int
// 	person
// }

// 8 - Nested Struct
type student struct {
	person struct { // 10 - tag property
		name string `tag1`
		age int `tag2`
	}
	grade int
	hobbies []string
}

// 9 - Horizontal Struct
// type person struct { name string; age int; hobbies []string}

// 11 - Type Alias
type Person struct {
	name string
	age int
}

type People = Person

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
	// var s1 = struct{
	// 	person
	// 	grade int
	// }{}

	// s1.person = person{"Hazlan", 23}
	// s1.grade = 100

	// fmt.Println("name  :", s1.person.name)
    // fmt.Println("age   :", s1.person.age)
    // fmt.Println("grade :", s1.grade)

	// fmt.Println("Name: ", s1.name)
	// fmt.Println("Age: ", s1.age)
	// fmt.Println("Age: ", s1.person.age)
	// fmt.Println("Grade: ", s1.grade)

	// var s2 = struct {
	// 	person
	// 	grade int
	// }{
	// 	person: person{"Rivano", 21},
	// 	grade: 95,
	// }

	// fmt.Println("Name: ", s2.name)
	// fmt.Println("Age: ", s2.age)
	// fmt.Println("Age: ", s2.person.age)
	// fmt.Println("Grade: ", s2.grade)

	// 6 - Slice and Struct Collab
	// var allStudents = []person{
	// 	{name: "Hazlan", age: 23},
	// 	{name: "Rivano", age: 21},
	// 	{name: "Riko", age: 21},
	// }

	// for _, student := range allStudents {
	// 	fmt.Println(student.name, " age is ", student.age)
	// }

	// 7 - Inisialisasi Slice Anonymous Struct
	// var allStudents = []struct{
	// 	person
	// 	grade int
	// }{
	// 	{person: person{"Hazlan", 23}, grade: 100},
	// 	{person: person{"Rivano", 21}, grade: 85},
	// 	{person: person{"Riko", 30}, grade: 100000},
	// }

	// for _, student := range allStudents {
	// 	fmt.Println(student.name)
	// 	fmt.Println(student.age)
	// 	fmt.Println(student.grade)
	// }

	// // 8 - Nested Struct (diatas)

	// // 9 - Struct Horizontal
	// var p1 = struct {name string, age int} {age: 21, name: "Hazlan"}
	// var p2 = struct {name string, age int} {"Rivano", 35}

	// 11 - Type Alias
	var p1 = Person{"Rivano", 22}
	fmt.Println(p1)
	
	var p2 = People{"Hazlan", 23}
	fmt.Println(p2)

	person := Person{"wick", 21}
	fmt.Println(People(person))
}