package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	temp, errT := template.ParseGlob("../pages/*.html")
	if errT != nil {
		println(errT)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "index", nil)
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "register", nil)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "login", nil)
	})

	http.HandleFunc("/process", processor)

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe("localhost:8080", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {

	fpseu := r.FormValue("pseudo")
	fmail := r.FormValue("adresse_mail")
	fname := r.FormValue("prenom")
	fsecname := r.FormValue("nom")
	fpassword := r.FormValue("password")

	fmt.Println(fpseu, fmail, fname, fsecname, fpassword)

	println(fmail, fpseu)

	type UserR struct {
		pseudo     string
		mail       string
		prenom     string
		secondname string
		password   string
	}

	d := UserR{fpseu, fmail, fname, fsecname, fpassword}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum")

	if err != nil {
		log.Fatal(err)
		defer db.Close()
	}

	fmt.Println(d.pseudo, d.secondname, d.prenom, d.mail, d.password)

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
