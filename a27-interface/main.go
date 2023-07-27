package main

import (
	f "fmt"
	m "math"
)

// 1 - Interface
// type hitung interface {
// 	luas() float64
// 	keliling() float64
// }

// // Lingkaran
// type lingkaran struct {
// 	diameter float64
// }

// func (l lingkaran) jariJari() float64 {
// 	return l.diameter / 2
// }

// func (l lingkaran) luas() float64 {
// 	return m.Pi * m.Pow(l.jariJari(), 2)
// }

// func (l lingkaran) keliling() float64 {
// 	return m.Pi * l.diameter
// }

// // Persegi
// type persegi struct {
// 	sisi float64
// }

// func (p persegi) luas() float64 {
// 	return m.Pow(p.sisi, 2)
// }

// func (p persegi) keliling() float64 {
// 	return 4 * p.sisi
// }

// 2 - Embedded Interface
type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return m.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return m.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func main() {
	// 1 - Interface
	// var bangunDatar hitung

	// bangunDatar = persegi{10.0}
	// f.Println("===== Persegi")
	// f.Println("Luas            :", bangunDatar.luas())
	// f.Println("Keliling        :", bangunDatar.keliling())
	
	// bangunDatar = lingkaran{14.0}
	// f.Println("===== Persegi")
	// f.Println("Luas            :", bangunDatar.luas())
	// f.Println("Keliling        :", bangunDatar.keliling())
	// f.Println("Jari-jari	   :", bangunDatar.(lingkaran).jariJari())

	// 2 - Embedded Intreface
	var hitungKubus hitung = &kubus{4}

	f.Println("======== Kubus")
	f.Println("Volume         :", hitungKubus.volume())
	f.Println("Luas           :", hitungKubus.luas())
	f.Println("Keliling       :", hitungKubus.keliling())

}