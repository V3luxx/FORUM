package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Selector(email string) Utilisateurs {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum")

	if err != nil {
		log.Fatal(err)
		defer db.Close()
	}
	var user Utilisateurs
	err0 := db.QueryRow("SELECT adresse_mail, mot_de_passe FROM utilisateurs where adresse_mail = ?", email).Scan(&user.Adresse_mail, &user.Mot_de_passe)
	if err0 != nil {
		println(err.Error())
	}
	return user
}
