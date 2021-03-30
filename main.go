package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"pioyi/golang_api/database"
	"pioyi/golang_api/utils"

	"github.com/joho/godotenv"
)

func main() {
	// Setting up the dotenv library
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file!")
	}

	// Setting up the databse
	database.Connect()

	// Closing it once the main function ends
	// (When the HTTP route listeners close due to a server shutdown)
	defer database.SQLDatabase.Close()

	// Initializing Routes
	router := utils.CreateRoutes()

	// Listening on port, displaying a colored message
	fmt.Printf("\033[32mServer listening on localhost%s!\033[0m\n", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))
}
