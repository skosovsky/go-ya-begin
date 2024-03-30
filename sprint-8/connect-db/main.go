package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "../db-demo/demo.db")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()
}
