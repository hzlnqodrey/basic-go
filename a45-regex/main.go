package main

import (
	"fmt"
	"regexp"
)

func main() {
	var str = "hazlan m qodri"

	// regexp.Compile()
	var regex, err = regexp.Compile(`[a-z]+`)
	if err != nil {
		fmt.Println(err.Error())
	}

	// .FindAllStrings()
	var STR1 = regex.FindAllString(str, 1)
	fmt.Println(STR1)

	var STR2 = regex.FindAllString(str, 2)
	fmt.Println(STR2)

	var STR3 = regex.FindAllString(str, -1)
	fmt.Println(STR3)

	// .MatchString()
	var text = "banana soup sapi"
	var isMatched, _ = regexp.MatchString(`[a-z]+`, text)
	fmt.Println(isMatched)
	var isMatched2, _ = regexp.MatchString(`[A-Z]+`, text)
	fmt.Println(isMatched2)

	// .FindString() - First String Found
	var text1 = "melon anjing kucing"
	var regexRes, _ = regexp.Compile(`[a-z]+`)

	var strr = regexRes.FindString(text1)
	fmt.Println(strr)

	// .FindStringIndex() - First String Found
	// hanya mengembalikan index string yg pertama kali ditemukan

	// .ReplaceAllString() - replace matched
	var text2 = "melon SAPI kucing ANJING"
	var regexKecil, _ = regexp.Compile(`[a-z]+`)
	var regexBesar, _ = regexp.Compile(`[A-Z]+`)

	var Kecil = regexKecil.ReplaceAllString(text2, "buah")
	fmt.Println(Kecil)

	var Besar = regexBesar.ReplaceAllString(text2, "hewan")
	fmt.Println(Besar)

	// .ReplaceAllStringFunc()
	var text3 = "banana sugar soup"
	var regex_, _ = regexp.Compile(`[a-z]+`)

	var regfunc = regex_.ReplaceAllStringFunc(text3, func(each string) string {
		if each == "sugar" {
			return "potato"
		}

		return each
	})

	fmt.Println(regfunc)

	// .Split()
	var nama = "Hazlan Muhammad Qodri"
	var regexxx, _ = regexp.Compile(`[a-b]+`)

	var str10 = regexxx.Split(nama, -1)
	fmt.Printf("%#v \n", str10)

}
