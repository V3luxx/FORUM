package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func login(w http.ResponseWriter, r *http.Request) {

	fpseu := r.FormValue("pseudo")
	fpassword := r.FormValue("password")

	fmt.Println(fpseu, fpassword)

	println(fpseu, fpassword)

	type UserR struct {
		pseudo   string
		password string
	}

	d := UserR{fpseu, fpassword}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum")

	if err != nil {
		log.Fatal(err)
		defer db.Close()
	}

	fmt.Println(d.pseudo, d.password)

	result, err := db.Exec("INSERT INTO utilisateurs (pseudo, nom, prenom, adresse_mail, mot_de_passe) VALUES (?,?,?,?,?)", d.pseudo, d.secondname, d.prenom, d.mail, d.password)
	Selector()
	fmt.Println("données rentrés")
	if err != nil {
		fmt.Errorf("addUser: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Print(id)
	}

	http.Redirect(w, r, "index", 301)
}
