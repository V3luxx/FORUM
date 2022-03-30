package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Insert() []Utilisateurs {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var user []Utilisateurs
	res, err := db.Query("INSERT INTO Utilisateurs Values(")
	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var oneUser Utilisateurs
		err := res.Scan(&oneUser.Id_utilisateurs, &oneUser.Nom, &oneUser.Pseudo, &oneUser.Prenom, &oneUser.Adresse_mail, &oneUser.Mot_de_passe)

		if err != nil {
			log.Fatal(err)
		}
		user = append(user, oneUser)
		/* fmt.Printf("%v\n", u) */
	}
	defer res.Close()
	return user
}
