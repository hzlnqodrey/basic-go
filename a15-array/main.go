package main

import "fmt"
import "strings"

func main() {
	
	var names[4] string

	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	for i := 0; i < len(names); i++ {
		fmt.Print(strings.ToUpper(names[i]), " ")
	}

}