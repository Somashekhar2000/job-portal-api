package database

import (
	"log"
	"project/internal/model"
)

// CreateTable initializes and updates database tables.
func CreateTable() {
	// Open the database connection
	db, err := Open()
	if err != nil {
		log.Fatalln("Failed to open the database:", err)
	}

	// Drop the "User" table if it exists
	if err = db.Migrator().DropTable(&model.User{}); err != nil {
		log.Fatalln("Failed to drop the 'User' table:", err)
	}

	// Create or update the "User" table
	if err = db.Migrator().AutoMigrate(&model.User{}); err != nil {
		log.Fatalln(err)
	}

}
