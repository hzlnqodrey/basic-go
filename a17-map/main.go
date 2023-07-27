package main

import "fmt"

func main() {
	
	var chicken = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}
	
	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}

	var chicken2 = map[string]int{
		"satu": 1,
		"dua": 2,
	}

	fmt.Println(len(chicken2))
	fmt.Println(chicken2)
	
	delete(chicken2, "satu")
	
	fmt.Println(len(chicken2))
	fmt.Println(chicken2)

	var chicken10 = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken10["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	var chickens = []map[string]string{
		{"name": "chicken blue",   "gender": "male"},
		{"name": "chicken red",    "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["name"], " \t", chicken["gender"])
	}

}