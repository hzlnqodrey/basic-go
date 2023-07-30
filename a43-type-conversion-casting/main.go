package main

import (
	"fmt"
	"strconv"
)

func main() {

	// strconv.atoi() => converting string to int
	var str = "12417823781231"
	var num, err = strconv.Atoi(str)
	if err == nil {
		fmt.Println(num)
	}

}