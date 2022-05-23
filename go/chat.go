package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)



func chater(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("msg")
	println(message)

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum")

	if err != nil {
		log.Fatal(err)
		defer db.Close()
	}

	chatmes, err := db.Exec("INSERT INTO stock_forum (commentaires) VALUES (?)", message)

	if err != nil {
		fmt.Println("addmsg: %v", err)
	}
	id, err := chatmes.LastInsertId()
	if err != nil {
		fmt.Print(id)
	}

	http.Redirect(w, r, "/chat", 301)
}
