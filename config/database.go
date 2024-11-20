package config

import (
	"database/sql"
	_ "github.com/lib/pq" // Import PostgreSQL driver
	"log"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error

	// Database connection string (adjust credentials as needed)
	dsn := "host=localhost user=your_username password=your_password dbname=farmer_market port=5432 sslmode=disable"
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		log.Fatalf("Database connection error: %v\n", err)
	}

	// Test query to validate the connection
	err = DB.QueryRow("SELECT 1").Scan(new(int))
	if err != nil {
		log.Fatalf("Test query failed: %v\n", err)
	}

	log.Println("Database connected successfully!")
}
