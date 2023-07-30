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

	// CASTING
	// Casting Literal (just use method of the type)
	var f3 float32 = float32(52.01)
	fmt.Println(f3)

	// Casting string to byte to string
	var s10 = "text"
	var b = []byte(s10)
	fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3])

	var byte1 = []byte{116, 101, 120, 116}
	var s11 = string(byte1)
	fmt.Printf("%s \n", s11)

	// Casting string <-> int
	var c int64 = int64('h')
	fmt.Println(c) // 104

	var d string = string(104)
	fmt.Println(d) // h

	// Type Assertions Pada Interface Kosong
	var data  = map[string]interface{}{
		"nama" : "Hazlan",
		"umur" : 123,
		"tinggi" : 174.1,
		"isMale" : true,
		"hobby" : []string{"coding", "reading"},
	}

	fmt.Println(data["nama"].(string))
	fmt.Println(data["umur"].(int))
	fmt.Println(data["tinggi"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobby"].([]string))

	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		case bool:
			fmt.Println(val.(bool))
		case []string:
			fmt.Println(val.([]string))
		default:
			fmt.Println(val.(int))
		}
	}
}
