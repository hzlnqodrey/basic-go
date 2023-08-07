package main

import "fmt"
import "flag"

func flagPrint() {
	var name = flag.String("name",  "Anonymous", "type your name")
	var age = flag.Int64("age", 25, "type your age")

	flag.Parse()
	fmt.Printf("name\t: %s\n", *name)
	fmt.Printf("age\t: %d\n", *age)
}
