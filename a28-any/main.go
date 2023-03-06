package main

import (
	"fmt"
)

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
}