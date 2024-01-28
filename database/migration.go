package database

import (
	"backend-technical-test/models"
	"backend-technical-test/pkg/postgres"
	"fmt"
	"sort"
	"strings"
)

func RunMigration() {
	// Migrate defined models
	err := postgres.DB.AutoMigrate(
		&models.List{},
		&models.Sublist{},
		&models.Migration{},
	)
	if err != nil {
		fmt.Println("Failed to migrate database")
		panic(err)
	}
	fmt.Println("Database Migrated!")

	// Find existing table names in the public schema
	var existingTables []string
	postgres.DB.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Scan(&existingTables)

	// Sort existing table names for easier comparison
	sort.Strings(existingTables)

	// Insert missing tables into the migration table
	for _, tableName := range existingTables {
		migrationName := fmt.Sprintf("create_%s_table", strings.ToLower(tableName))
		if !isTableMigrated(migrationName) {
			err := insertMigrationRecord(migrationName)
			if err != nil {
				fmt.Printf("Failed to insert migration record for table %s: %v\n", tableName, err)
			}
		}
	}
}

func isTableMigrated(migrationName string) bool {
	var count int64
	postgres.DB.Raw("SELECT COUNT(*) FROM migrations WHERE name = ?", migrationName).Scan(&count)
	return count > 0
}

func insertMigrationRecord(migrationName string) error {
	result := postgres.DB.Exec("INSERT INTO migrations (name, created_at) VALUES (?, NOW())", migrationName)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
