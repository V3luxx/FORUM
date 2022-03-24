package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func maine() {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Query("SELECT * FROM utilisateurs;")
	res, err := db.Query("INSERT INTO ")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {

		var user Utilisateurs
		err := res.Scan(&user.Id_utilisateurs, &user.Nom, &user.Pseudo, &user.Prenom, &user.Adresse_mail)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", user)
	}
}
