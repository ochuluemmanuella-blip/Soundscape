package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connstr := "postgresql://neondb_owner:npg_lGzB19NyxUbE@ep-bold-glitter-a56e8n3k-pooler.us-east-2.aws.neon.tech/neondb?sslmode=require&channel_binding=require"
	db, err := sql.Open("postgres", connstr)

	if err != nil {
		log.Fatalf("Error opening database definition: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("cannot connect to Postgres: %v", err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")
	schema := `  
	CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	email TEXT UNIQUE NOT NULL
	);`
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	insertQuery := "INSERT INTO users (name, email) VALUES ($1, $2) ON CONFLICT DO NOTHING"
	_, err = db.Exec(insertQuery, "Test User", "test@example.com")
	if err != nil {
		log.Fatalf("Failed to insert row: %v", err)
	}

	var fetchedID int
	var fetchedEmail string
	searchEmail := "test@example.com"

	selectQuery := "SELECT id, email FROM users WHERE email = $1"
	err = db.QueryRow(selectQuery, searchEmail).Scan(&fetchedID, &fetchedEmail)
	if err == sql.ErrNoRows {
		log.Printf("No user found with email:  %s\n", searchEmail)
	} else if err != nil {
		log.Fatalf("Query failed: %v", err)
	} else {
		fmt.Printf("Found user in database ->  ID: %d, Email:  %s\n", fetchedID, fetchedEmail)
	}
}
