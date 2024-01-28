package database

import (
	"backend-technical-test/models"
	"backend-technical-test/pkg/postgres"
	"fmt"
)

func RunMigration() {
	err := postgres.DB.AutoMigrate(
		&models.List{},
		&models.Sublist{},
	)

	if err != nil {
		fmt.Println("Failed to migrate database")
		panic(err)
	}

	fmt.Println("Database Migrated!")
}