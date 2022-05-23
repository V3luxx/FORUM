package main

import (
	"database/sql"
	"log"
)

func getMessage() []string {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum")

	if err != nil {
		log.Fatal(err)
		defer db.Close()
	}

	getMessages, err := db.Query("SELECT commentaires FROM `stock_forum`")
	if err != nil {
		println(err.Error())
	}

	var allMessage []string

	for getMessages.Next() {
		var message string
		err = getMessages.Scan(&message)
		if err != nil {
			println(err.Error())
			return allMessage
		}
		allMessage = append(allMessage, message)
	}
	return allMessage
}
