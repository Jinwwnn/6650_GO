/*
 * Album Store API
 *
 * CS6650 Fall 2023
 *
 * API version: 1.0.0
 * Contact: i.gorton@northeasern.edu
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	sw "6650_GO/go"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Printf("Server started")

	// Get database connection string from environment variable
	dbDSN := os.Getenv("DB_DSN")
	if dbDSN == "" {
		// Default DSN if not set
		dbDSN = "mydbuser:mydbpass123@tcp(demo-mysql.c7uus8adixh0.us-west-2.rds.amazonaws.com:3306)/mydemodb"
	}

	// Initialize database connection
	db, err := sql.Open("mysql", dbDSN)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Test database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Printf("Successfully connected to database")

	// Initialize database in the API package
	sw.InitDB(db)

	// Initialize router
	router := sw.NewRouter()

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
}
