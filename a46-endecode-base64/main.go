// Go memiliki package encoding/base64,
// berisikan fungsi-fungsi untuk kebutuhan encode dan decode data ke base64 dan sebaliknya. 
// Data yang akan di-encode harus bertipe []byte, 
// perlu dilakukan casting untuk data-data yang belum sesuai tipenya.

package main

import(
	"encoding/base64"
	"fmt"
)

func main() {
	// EncodeToString() and DecodeString()
	var data = "Hazlan M Qodri"
	fmt.Println("data: ", data)
	
	// encode
	var EncodeByte = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded: ", EncodeByte)

	// decode
	var DecodeByte, _ = base64.StdEncoding.DecodeString(EncodeByte)
	var DecodeString = string(DecodeByte)
	fmt.Println("decoded: ", DecodeString)
	
}