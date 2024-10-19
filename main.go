package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	// Database connection parameters
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "password"
	dbName := "postgres"

	// Build the connection string
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Ensure the schema_migrations table exists
	err = createMigrationsTable(db)
	if err != nil {
		log.Fatalf("Failed to create migrations table: %v", err)
	}

	// Get the list of applied migrations
	appliedMigrations, err := getAppliedMigrations(db)
	if err != nil {
		log.Fatalf("Failed to get applied migrations: %v", err)
	}

	// Read migration files
	migrationFiles, err := ioutil.ReadDir("migrations")
	if err != nil {
		log.Fatalf("Failed to read migrations directory: %v", err)
	}

	// Filter and sort migration files
	migrations := filterAndSortMigrations(migrationFiles)

	// Apply pending migrations
	for _, migration := range migrations {
		if _, applied := appliedMigrations[migration]; applied {
			fmt.Printf("Skipping already applied migration: %s\n", migration)
			continue
		}

		err = applyMigration(db, migration)
		if err != nil {
			log.Fatalf("Failed to apply migration %s: %v", migration, err)
		}

		fmt.Printf("Applied migration: %s\n", migration)
	}

	fmt.Println("All migrations applied successfully.")
}

// createMigrationsTable ensures the schema_migrations table exists
func createMigrationsTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS schema_migrations (
        version VARCHAR(255) PRIMARY KEY
    );
    `
	_, err := db.Exec(query)
	return err
}

// getAppliedMigrations retrieves the list of applied migrations
func getAppliedMigrations(db *sql.DB) (map[string]bool, error) {
	rows, err := db.Query("SELECT version FROM schema_migrations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	migrations := make(map[string]bool)
	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			return nil, err
		}
		migrations[version] = true
	}
	return migrations, nil
}

// filterAndSortMigrations filters .sql files and sorts them
func filterAndSortMigrations(files []os.FileInfo) []string {
	var migrations []string
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			migrations = append(migrations, file.Name())
		}
	}
	sort.Strings(migrations)
	return migrations
}

// applyMigration reads and executes a migration file
func applyMigration(db *sql.DB, filename string) error {
	path := filepath.Join("migrations", filename)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Execute the SQL commands
	_, err = db.Exec(string(content))
	if err != nil {
		return err
	}

	// Record the migration as applied
	_, err = db.Exec(
		"INSERT INTO schema_migrations (version) VALUES ($1)",
		filename,
	)
	return err
}
