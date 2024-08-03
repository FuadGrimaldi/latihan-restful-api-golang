package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetConnection(dbname string)(*sql.DB, error) {
	// init SQLite3 database
	db, err := sql.Open("sqlite3", dbname) 
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to the database: %s", err)
	}
	fmt.Println("Connected successfully to the database")
	return db, nil
}
func ConnTest() {
	dsn := "todo.db"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table';")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Tables:")
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
	}
}