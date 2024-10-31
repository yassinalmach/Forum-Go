package main

import (
	"forum/database"
	"log"
)

func main() {
	err := database.SetupDatabase()
	if err != nil {
		log.Fatalln(err)
	}
	defer database.DB.Close()
}
