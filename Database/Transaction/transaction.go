package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/gocourse")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into users(id, name) values(?, ?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")
	_, err = stmt.Exec(1, "Thiago") // duplicate key

	if err != nil {
		tx.Rollback()  // don't execute any sql commands
		log.Fatal(err) // break script
	}

	tx.Commit() // execute all Exec()
}
