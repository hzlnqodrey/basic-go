package main

import "fmt"

func main() {

	// menggunakan var, tanpa tipe data, menggunakan perantara "="
	var firstName = "john"

	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	lastName := "wick"

	// deklarasi multi variabel
	var first, second, third string

	first, second, third = "satu", "dua", "tiga"

	// or

	var fourth, fifth, sixth string = "empat", "lima", "enam"

	// or | Dengan menggunakan teknik type inference, 
	// deklarasi multi variabel bisa dilakukan untuk variabel-variabel yang tipe data satu sama lainnya berbeda.

	seventh, eighth, ninth := "tujuh", 8, "sembilan"

	fmt.Printf("Hello, Sir %s %s!\n", firstName, lastName)

	fmt.Println(first, second, third)
	fmt.Println(fourth, fifth, sixth)
	fmt.Println(seventh, eighth, ninth)

	// variable underscore (predefined)
	/*
	Biasanya variabel underscore sering dimanfaatkan untuk menampung nilai balik fungsi yang tidak digunakan.
	*/
	name, _ := "i can see shit", "we wewewew"

	fmt.Println(name)
	
	// new pointer variable

	namaaa := new(string)

	fmt.Println(namaaa)
	fmt.Println(*namaaa)
	*namaaa = "nama pointer"
	fmt.Println(*namaaa)
}