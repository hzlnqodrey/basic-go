package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "http://kalipare.com/hello?name=Hazlan M Qodri&age=23"
	var u, err = url.Parse(urlString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)
	fmt.Printf("protocl: %s\n", u.Scheme) // http
	fmt.Printf("host: %s\n", u.Host) // kalipare.com:80
	fmt.Printf("path: %s\n", u.Path) // hello

	var name = u.Query()["name"][0] // Hazlan M Qodri
	var age = u.Query()["age"][0] // 23
	fmt.Printf("name %s. age: %s\n", name, age)
	
}