package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	// strconv.atoi() => converting string to int
	var str = "12417823781231"
	var num, err = strconv.Atoi(str)
	if err == nil {
		fmt.Println(num)
	}

	// strconv.Itoa() => converting int to string
	var namba = 890

	var str_2 = strconv.Itoa(namba)
	fmt.Println(str_2)

	n := reflect.TypeOf(num)
	r := reflect.TypeOf(str_2)
	fmt.Println(n, r)
}
