package swagger

import (
	"database/sql"
	"log"
)

// InitDB initializes the database connection
func InitDB(database *sql.DB) {
	db = database
	// Create table if not exists
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS albums (
			album_id VARCHAR(255) PRIMARY KEY,
			artist VARCHAR(255),
			title VARCHAR(255),
			year VARCHAR(4),
			image_data MEDIUMBLOB,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}
