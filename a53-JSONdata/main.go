package main

import "fmt"
import "encoding/json"

type User struct {
	FullName string `json:"Name"`
	Age int
}

func main() {
	var jsonString = `{"Name": "Hazlan M Qodri", "Age": 23}`
	var jsonData = []byte(jsonString)

	var data User

	// Fungsi unmarshal hanya menerima data json dalam bentuk  []byte
	// Casting string -> byte
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data.FullName)
	fmt.Println("age :", data.Age)

	// === 53.2. Decode JSON Ke map[string]interface{}  &  interface{}

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data1["Name"])
	fmt.Println("age: ", data1["Age"])

	// === 53.3 Decode Array JSON Ke Array Objek

	var jsonStringArray = `[
		{"Name": "Hazlan M Qodri", "Age": 23},
		{"Name": "Dumby Dumb", "Age": 32}
	]`

	var dataUser []User

	var err1 = json.Unmarshal([]byte(jsonStringArray), &dataUser)
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}

	fmt.Println("User 1 :", dataUser[0].FullName, " - ", dataUser[0].Age)
	fmt.Println("User 2 :", dataUser[1].FullName, " - ", dataUser[1].Age)
}