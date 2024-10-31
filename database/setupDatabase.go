package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func SetupDatabase() error {
	// create database, open creates a sql.DB object that acts as a connection pool.
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return fmt.Errorf("error create database.db: %v", err)
	}

	// check if the connection establish
	if err := db.Ping(); err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	// read schema design database from .sql file
	schema, err := os.ReadFile("./database/schema.sql")
	if err != nil {
		return fmt.Errorf("error reading schema.sql: %v", err)
	}

	// creating tables
	if _, err := db.Exec(string(schema)); err != nil {
		return fmt.Errorf("error executing schema.sql: %v", err)
	}

	DB = db
	return nil
}
