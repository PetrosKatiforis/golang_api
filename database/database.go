package database

import (
	"fmt"
	"log"
	"os"
	"pioyi/golang_api/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var SQLDatabase *gorm.DB

func Connect() {
	// Loading Database Configurations
	dbConnectionString := os.Getenv("DATABASE_CONNECTION_STRING")
	dialect := os.Getenv("DATABASE_DIALECT")

	db, err := gorm.Open(dialect, dbConnectionString)
	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	} else {
		fmt.Println("\nDatabase connected successfully!")
	}

	// Migrating tables
	db.AutoMigrate(&entity.User{}, &entity.Post{})

	// Exporting the database instance
	SQLDatabase = db
}
