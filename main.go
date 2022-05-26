package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mahendrafathan/go-api-boilerplate/helpers"
	"github.com/mahendrafathan/go-api-boilerplate/models"
	"github.com/mahendrafathan/go-api-boilerplate/router"

	"github.com/joho/godotenv"
)

func init() {
	if envLoadError := godotenv.Load(); envLoadError != nil {
		log.Fatal("[ ERROR ] Failed to load .env file")
	}
}

func main() {
	var PORT string
	db := helpers.CreateDatabaseInstance()
	if migrateError := db.AutoMigrate(&models.User{}, &models.Book{}); migrateError != nil {
		log.Fatal("[ ERROR ] Couldn't migrate models!")
	}

	router := router.RegisterRoutes(db)
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "9090"
	}

	fmt.Printf("[ OK ] Server is Started and Listening on port: %v", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
