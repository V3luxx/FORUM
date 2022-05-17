package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var temp *template.Template

/* func init() {
	temp, errT := template.ParseGlob("forum/*.html")
	if errT != nil {
		println(errT)
	}
} */

func main() {
	var errT error
	temp, errT = template.ParseGlob("./*.html")
	if errT != nil {
		println(errT)
	}
	http.HandleFunc("./", index)
	http.HandleFunc("/process", processor)
	http.ListenAndServe("localhost:8080", nil)

}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "Post" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	fpseu := r.FormValue("pseudo")
	fmail := r.FormValue("mail")
	fname := r.FormValue("name")
	fsecname := r.FormValue("secondname")
	fpassword := r.FormValue("password")

	println(fmail)

	d := struct {
		pseudo     string
		mail       string
		prenom     string
		secondname string
		password   string
	}{
		pseudo:     fpseu,
		mail:       fmail,
		prenom:     fname,
		secondname: fsecname,
		password:   fpassword,
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum")

	if err != nil {
		log.Fatal(err)
		defer db.Close()
	}

	result, err := db.Exec("INSERT INTO users (pseudo, nom, prenom, adresse_mail, mot_de_passe) VALUES (?, ?, ?, ?, ?, ?, ?)", &d.pseudo, &d.secondname, &d.prenom, &d.mail, &d.password)
	Selector()
	if err != nil {
		fmt.Errorf("addUser: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Errorf("addUser: %v", err, id)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("fonctionnel")
	temp.ExecuteTemplate(w, "./register.html", nil)
}
