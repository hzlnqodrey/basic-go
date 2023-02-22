package main

import (
	"fmt"
	"strings"
)

type student struct {
	grade int
	name string
}

// Method Declaration

func (s student) sayHello() {
	fmt.Println("halo ", s.name)
}

func (s student) getNickname(i int) string {
	var Nickname = strings.Split(s.name, " ")[i - 1]
	return Nickname
}

func main() {
	var s1 = student{90, "john wick"}
	s1.sayHello()

	var name = s1.getNickname(2)
	fmt.Printf("nama panggilan: %s", name)
}