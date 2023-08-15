// Web Service API adalah sebuah web yang menerima request dari client dan menghasilkan response, biasa berupa JSON/XML
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	ID    string
	Name  string
	Grade int
}

var data = []student{
	student{"E001", "ethan", 21},
	student{"H001", "hazlan", 23},
	student{"R001", "rivano", 21},
	student{"R002", "riko", 21},
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data) // parse json
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// Method  r.FormValue()  digunakan untuk mengambil data form yang dikirim dari client, 
		// pada konteks ini data yang dimaksud adalah ID
		var ID = r.FormValue("ID")
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == ID {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}

		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("Starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}