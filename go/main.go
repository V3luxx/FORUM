package main

import (
	"net/http"
)

var u1 []Utilisateurs

func main() {
	u1 = Selector()
	for _, test := range u1 {
		println(test.Id_utilisateurs)
	}
	println("end")
	http.ListenAndServe("localhost:8080", nil)
}
