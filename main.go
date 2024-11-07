package main

import (
	"fmt"
	"forum/database"
	"forum/handlers/api"
	"log"
	"net/http"
)

func main() {
	err := database.SetupDatabase()
	if err != nil {
		log.Fatalln(err)
	}
	defer database.DB.Close()

	mux := http.NewServeMux()
	setupRoutes(mux)

	fmt.Println("server running on: http://localhost:2000")
	if err := http.ListenAndServe(":2000", mux); err != nil {
		log.Fatalln("running the server failed")
	}
}

func setupRoutes(mux *http.ServeMux) {
	// login routes
	mux.HandleFunc("POST /api/register", api.RegisterUser)
	mux.HandleFunc("POST /api/login", api.LoginUser)
}
