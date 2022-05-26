package helpers

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateDatabaseInstance() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("[ ERROR ] Unable to connect with mysql!\n", err)
	}

	fmt.Println("[ OK ] Connected to the DB!")

	return db
}
