package main

import (
	"html/template"
	"net/http"
)

var u1 []Utilisateurs

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("forum/*.html"))
}

func main() {
	u1 = Selector()
	for _, test := range u1 {
		println(test.Id_utilisateurs)
	}
	println("test")
	http.HandleFunc("/", index)
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

	d := struct {
		pseudo     string
		mail       string
		name       string
		secondname string
	}{
		pseudo:     fpseu,
		mail:       fmail,
		name:       fname,
		secondname: fsecname,
	}
	temp.ExecuteTemplate(w, "processor.html", d)
}

func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "register.html", nil)
}
