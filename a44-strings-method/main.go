package main

import (
	"fmt"
	"strings"
)

func main() {

	// Strings.Contains()
	var isSame = strings.Contains("john wick", "wick")
	fmt.Println(isSame)

	// Strings.HasPrefix()
	var isPrefix = strings.HasPrefix("Hazlan Qodrey", "Haz")
	fmt.Println(isPrefix)

	var isPrefix2 = strings.HasPrefix("Hazlan Qodrey", "rey")
	fmt.Println(isPrefix2)
	
	// Strings.HasSuffix()
	var isSuffix = strings.HasSuffix("Hazlan Qodrey", "rey")
	fmt.Println(isSuffix)

	// Strings.Count()
	var sCount = strings.Count("Hazlan M Qodri", "a") // 2
	fmt.Println(sCount)

	// Strings.Index()
	var index1 = strings.Index("Hazlan M Qodri", "az") // 9
	fmt.Println(index1)

	// Strings.Replace()
	var text = "banana"
	var find = "a"
	var with = "o"

	var newText = strings.Replace(text, find, with, -1)
	fmt.Println(newText)
	
	var newText2 = strings.Replace(text, find, with, 1)
	fmt.Println(newText2)

	// strings.Repeat()
	var str = strings.Repeat("naruto | ", 4)
	fmt.Println(str)

	// strings.Split()
	var strSplit = strings.Split("Hazlan Muhammad Qodri", " ")
	fmt.Println(strSplit[0])
	fmt.Println(strSplit[1])
	fmt.Println(strSplit[2])

	// strings.Join()
	var strJoin = strings.Join([]string{"Hazlan", "Muhammad", "Qodri"}, " - ")
	fmt.Println(strJoin)

	// strings.ToLower()
	var strLower = strings.ToLower("aLAy")
	fmt.Println(strLower) // "alay"

	// strings.ToUpper()
	var strUpper = strings.ToUpper("anjing")
	fmt.Println(strUpper)
}