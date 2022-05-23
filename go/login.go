package main

import "net/http"

func testStr(str1, str2 string) bool {
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return false
		}
	}

	return true
}

func login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	userCheck := Selector(email)
	if testStr(email, userCheck.Adresse_mail) && testStr(password, userCheck.Mot_de_passe) {
		http.Redirect(w, r, "index", 301)
		println("cest ok l'ami !")
	} else {
		http.Redirect(w, r, "/login", 301)
	}
}
