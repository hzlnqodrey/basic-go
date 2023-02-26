package main

import (
	f "fmt"
	m "math"
)


type hitung interface {
	luas() float64
	keliling() float64
}

// Lingkaran
type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return m.Pi * m.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return m.Pi * l.diameter
}

// Persegi
type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return m.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return 4 * p.sisi
}

func main() {
	var bangunDatar hitung

	bangunDatar = persegi{10.0}
	f.Println("===== Persegi")
	f.Println("Luas            :", bangunDatar.luas())
	f.Println("Keliling        :", bangunDatar.keliling())
	
	bangunDatar = lingkaran{14.0}
	f.Println("===== Persegi")
	f.Println("Luas            :", bangunDatar.luas())
	f.Println("Keliling        :", bangunDatar.keliling())
	f.Println("Jari-jari	   :", bangunDatar.(lingkaran).jariJari())

}