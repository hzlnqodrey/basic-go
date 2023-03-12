package main

import (
	"fmt"
)

type student struct {
	name string
	grade int
}

// 1 -  Method Declaration
// func (s student) sayHello() {
// 	fmt.Println("halo ", s.name)
// }

// func (s student) getNickname(i int) string {
// 	var Nickname = strings.Split(s.name, " ")[i - 1]
// 	return Nickname
// }

// 2 - Method Pointer
func (s student) changeName1(name string) {
	fmt.Println("----> on changename1, name changed to ", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("----> on changename2, name changed to ", name)
	s.name = name
}

func main() {
	// var s1 = student{90, "john wick"}
	// s1.sayHello()

	// var name = s1.getNickname(2)
	// fmt.Printf("nama panggilan: %s", name)

	// 2 - Method Pointer
	var s1 = student{"Hazlan Qodrey", 21}
	fmt.Println("s1 before", s1.name)

	s1.changeName1("Rivano ATK")
	fmt.Println("s1 after changename1", s1.name)
	
	s1.changeName2("Riko A")
	fmt.Println("s1 after changename2", s1.name)

}