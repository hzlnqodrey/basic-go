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

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data1["Name"])
	fmt.Println("age: ", data1["Age"])
}