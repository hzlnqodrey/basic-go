package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age int
}

func main() {

	var secret interface{}

	secret = "Hazlan Muhammad Qodri"
	fmt.Println(secret)

	secret = []string{"Banana", "Melon"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	// ===

	var data map[string]interface{}

	data = map[string]interface{}{
		"Nama"		: "Hazlan Muhammad Qodri",
		"Grade"		: 90,
		"Breakfast"	: []string{"Nasi Uduk", "Lontong"},
	} 
	
	fmt.Println(data["Nama"], data["Grade"], data["Breakfast"])

	// === Any [Type ALIAS]
	var data2 map[string]any

	data2 = map[string]any{
		"Nama"		: "Hazlan Muhammad Qodri",
		"Grade"		: 90,
		"Breakfast"	: []string{"Nasi Uduk", "Lontong"},		
	}

	fmt.Println(data2["Nama"], data2["Grade"], data2["Breakfast"])

	// Casting Variable Any

	var benda interface{}

	// Casting to INT
	benda = 2
	var multipled int = benda.(int) * 10
	fmt.Println("2 dikali sepuluh sama dengan", multipled)

	// Casting to Array of String
	benda = []string{"Kursi", "Meja", "PC"}
	things := strings.Join(benda.([]string), ", ")
	fmt.Println(things, " are my favorite things")

	// === Interface Casting Object
	var secret2 interface{} = &person{"Hazlan", 23}
	
	var name = secret2.(*person).name
	fmt.Println(name)
}