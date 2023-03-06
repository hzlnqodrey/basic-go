package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Introduction Reflect - ReflectValueOf() dan Reflect.Type()
	var number int = 23

	var reflectValue = reflect.ValueOf(number)

	fmt.Println("Tipe data dari reflectValue adalah", reflectValue.Type())
	
	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai Variabel :", reflectValue.Int())
	}

	// If value just want to output not passing, just use .Interface()
	var number2 int = 100

	var reflectValue2 = reflect.ValueOf(number2)
	
	fmt.Println("Tipe data dari reflectValue adalah", reflectValue2.Type())
	fmt.Println("Nilai dari reflectValue adalah", reflectValue2.Interface())

}