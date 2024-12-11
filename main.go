package main

import (
	"database/sql"
	"flag"
	"log"
	"project-manager/config"
	"project-manager/db"
)

// Handles database migrations.
func runMigrations(d *sql.DB) {
	log.Println("Starting database migrations...")
	if err := db.MigrateTables(d, 10); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Database migrations completed successfully.")
}

func main() {
	// Parse command-line flags
	migrateFlag := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	// Connect to the database
	dbConn, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if err := dbConn.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	// Run migrations if the flag is set
	if *migrateFlag {
		runMigrations(dbConn)
	}
}

