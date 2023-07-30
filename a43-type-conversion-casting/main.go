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

	// strconv.ParseInt() => String number to number
	var str3 = "124"
	var num3, err3 = strconv.ParseInt(str3, 10, 64)
	if err3 == nil {
		fmt.Println(num3)
	}

	var str4 = "1010"
	var num4, err4 = strconv.ParseInt(str4, 2, 8)

	if err4 == nil {
		fmt.Println(num4) // 10
	}	

	// strconv.FormatInt => number to string number
	var num5 = int64(990)
	var str5 = strconv.FormatInt(num5, 10)
	if err == nil {
		fmt.Println(str5)
	}

	// strconv.ParseFloat() => string float to float
	var str6 = "20.140"
	var float1, err6 = strconv.ParseFloat(str6, 32)
	if err6 == nil {
		fmt.Println(float1)
	}

	// strconv.FormatFloat() => float to string flaot
	var float2 = 55.92
	var s7 = strconv.FormatFloat(float2, 'f', 6, 64)
	fmt.Println(s7)

	// strconv.ParseBool() => string to bool
	var s8 = "true"
	var b1, err7 = strconv.ParseBool(s8)
	if err7 == nil {
		fmt.Println(b1)
	}

	// strconv.FormatBool() => bool to string
	var b2 = false
	var s9 = strconv.FormatBool(b2)
	fmt.Println(s9)
}
