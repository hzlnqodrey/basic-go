package main

import (
	"fmt"
	"reflect"
)

// 3 - Reflect Access Object Property
type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var refVal = reflect.ValueOf(s)

	if refVal.Kind() == reflect.Ptr {
		refVal = refVal.Elem()
	}

	var refType = refVal.Type()

	for i := 0; i < refVal.NumField(); i++ {
		fmt.Println("Nama		: ", refType.Field(i).Name)
		fmt.Println("Tipe Data		: ", refType.Field(i).Type)
		fmt.Println("Nilai		: ", refVal.Field(i).Interface())
		fmt.Println(" ")
	}
}

// 4 - Reflect get Info from Object Method
func (s *student) SetName(name string) {
	s.Name = name
}

func main() {

	// 3 - Reflect Access Object Property
	var s1 = &student{Name: "Hazlan", Grade: 85}
	s1.getPropertyInfo()
	fmt.Println("Nama sblm: ", s1.Name)

	// 4 - Reflect get Info from Object Method
	var studRefVal = reflect.ValueOf(s1)
	var methodStudent = studRefVal.MethodByName("SetName")

	methodStudent.Call([]reflect.Value{
		reflect.ValueOf("Hazlan Qodrey"),
	})

	fmt.Println("Nama ssdh: ", s1.Name)
	fmt.Println(" ")

	// 1 -  Introduction Reflect - ReflectValueOf() dan Reflect.Type()
	var number int = 23

	var reflectValue = reflect.ValueOf(number)

	fmt.Println("Tipe data dari reflectValue adalah", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai Variabel :", reflectValue.Int())
	}

	// 2 - If value just want to output not passing, just use .Interface()
	var number2 int = 100

	var reflectValue2 = reflect.ValueOf(number2)

	fmt.Println("Tipe data dari reflectValue adalah", reflectValue2.Type())
	fmt.Println("Nilai dari reflectValue adalah", reflectValue2.Interface())

	// other way to access real value of reflect variable is
	var Nilai = reflectValue2.Interface().(int)
	fmt.Println(Nilai * 100)

	// 3 - Reflect Access Object Property
}
