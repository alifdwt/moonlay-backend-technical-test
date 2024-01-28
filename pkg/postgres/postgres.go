package postgres

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	errr := godotenv.Load()
	if errr != nil {
		panic("Failed to load .env file!")
	}

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		"postgres",
		os.Getenv("DB_PORT"),
		os.Getenv("APP_TIMEZONE"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	dbName := os.Getenv("DB_NAME")
	checkAndCreateDatabase(DB, dbName)

	dsnNew := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		dbName,
		os.Getenv("DB_PORT"),
		os.Getenv("APP_TIMEZONE"),
	)

	DB, err = gorm.Open(postgres.Open(dsnNew), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connection Opened to Database")
}

func TestDatabase(host string, user string, password string, dbname string, port string, timezone string) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s",
		host,
		user,
		password,
		"postgres",
		port,
		timezone,
	)

	DB2, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	checkAndCreateDatabase(DB2, dbname)

	dsnNew := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s",
		host,
		user,
		password,
		dbname,
		port,
		timezone,
	)

	DB, err = gorm.Open(postgres.Open(dsnNew), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connection Opened to Database")
}

func checkAndCreateDatabase(DB *gorm.DB, dbName string) {
	var count int64
	result := DB.Raw("SELECT count(*) FROM pg_database WHERE datname = ?", dbName).Scan(&count)
	if result.Error != nil {
		panic(result.Error)
	}

	if count == 0 {
		createDatabase(DB, dbName)
	} else {
		fmt.Printf("Database %s already exists\n", dbName)
	}
}

func createDatabase(DB *gorm.DB, dbName string) {
	result := DB.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	// DB.Exec(fmt.Sprintf("INSERT INTO migrations (id, name, created_at) VALUES (1, 'create_database_%s', NOW())", dbName))
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Printf("Database %s created\n", dbName)
}
