package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Selector() []Utilisateurs {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum")

	if err != nil {
		log.Fatal(err)
		defer db.Close()
	}
	var user []Utilisateurs
	res, err := db.Query("SELECT pseudo, mot_de_passe FROM utilisateurs")
	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var oneUser Utilisateurs
		err := res.Scan(&oneUser.Pseudo, &oneUser.Mot_de_passe)

		if err != nil {
			log.Fatal(err)
		}
		user = append(user, oneUser)
		/* fmt.Printf("%v\n", u) */
		
	}
	defer res.Close()
	return user
}
