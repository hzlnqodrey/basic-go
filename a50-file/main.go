package main

import (
	"fmt"
	"os"
)

var path = "C:\\.qodri-and-his-cloud-adventure\\.backend-language\\golang-projects\\basic-go\\a50-file\\test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

// func createFile() {
// 	// deteksi apakah file sudah ada
// 	var _, err = os.Stat(path)

// 	// buat file baru jika belum ada
// 	if os.IsNotExist(err) {
// 		var file, err = os.Create(path)
// 		if isError(err) { return }
// 		defer file.Close()
// 	}

// 	fmt.Println("===> file berhasil dibuat", path)
// }

// edit file = buka akses write
func writeFile() {
	// buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) { return }
	defer file.Close()

	// tulis data ke file
	_, err = file.WriteString("hello world\n")
	if isError(err) { return }
	
	_, err = file.WriteString("Bab A50 - File Golang Basic Tutorial\n")
	if isError(err) { return }

	// save changes
	err = file.Sync()
	if isError(err) { return }

	fmt.Println("===> file berhasil ditulis")
}

func main() {
	// createFile()
	writeFile()
}