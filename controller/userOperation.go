package controller

import (
	"database/sql"
	"fmt"
	"forum/database"

	"golang.org/x/crypto/bcrypt"
)

func AddUser(email, username, password string) error {
	// check if data provided exists
	if email == "" || username == "" || password == "" {
		return fmt.Errorf("email, username and password are required")
	}

	// check if user already registred
	var isExist bool
	if err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ? OR username = ?)",
		email, username).Scan(&isExist); err != nil {
		return err
	} else if isExist {
		return fmt.Errorf("user already exist")
	}

	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// insert data
	if _, err := database.DB.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		username, email, string(hashedPass)); err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	return nil
}

func GetUser(username, password string) error {
	var hashedPassword string
	// check if username already exist
	err := database.DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		return fmt.Errorf("user not found: %v", err)
	} else if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return fmt.Errorf("incorrect password: %v", err)
	}

	return nil
}
