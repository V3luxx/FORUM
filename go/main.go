package main

import(
	"net/http"
)

func main() {
	u1 := Utilisateurs{}
	
	println(u1.Id_utilisateurs)

	http.ListenAndServe("localhost:8080", nil)
}

	