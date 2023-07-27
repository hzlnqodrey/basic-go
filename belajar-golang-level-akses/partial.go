package main

import f "fmt"

// go run main.go partial.go ---> run this in CLI when you want to execute main.go
// or
// FOR WINDOWS: [go run .]

func SayHello(name string) {
	f.Println("halo", name)
}