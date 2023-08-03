package main

// Hash termasuk one-way encryption, 
// membuat hasil dari hash tidak bisa dikembalikan ke text asli.

import(
	"fmt"
	"crypto/sha1"
)

func main() {
	var text = "this is secret"
	fmt.Println("Original data: ", text)
	
	var sha = sha1.New()
	sha.Write([]byte(text))
	var encrypted = sha.Sum(nil)

	var encryptedString = fmt.Sprintf("%x", encrypted)
	fmt.Println("hashed: ", encryptedString)
}