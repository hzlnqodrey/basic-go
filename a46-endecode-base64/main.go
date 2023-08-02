// Go memiliki package encoding/base64,
// berisikan fungsi-fungsi untuk kebutuhan encode dan decode data ke base64 dan sebaliknya.
// Data yang akan di-encode harus bertipe []byte,
// perlu dilakukan casting untuk data-data yang belum sesuai tipenya.

package main

import (
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

	// Encode() and Decode()
	var data2 = "Hazlan Qodrey"
	fmt.Println("\n data2:", data2)

	// Encode()
	var encoding = make([]byte, base64.StdEncoding.EncodedLen(len(data2)))
	base64.StdEncoding.Encode(encoding, []byte(data2))
	var EncodingString = string(encoding)
	fmt.Println("Encoding: ", EncodingString)

	// Decode()
	var decoding = make([]byte, base64.StdEncoding.DecodedLen(len(encoding)))
	var _, err  = base64.StdEncoding.Decode(decoding, encoding)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Decoding", string(decoding))

}	
