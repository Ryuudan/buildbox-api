package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import the pq driver
)

func PostgresConnect() (*sql.DB, error) {
	dbHost := "localhost" // Since your PostgreSQL container is running on the same host
	dbPort := "5435"      // The mapped port from your Docker Compose file
	dbName := "buildbox"
	dbUser := "bb_user"
	dbPassword := "bb_password"

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Attempt to ping the database to check the connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Successfully connected to the database
	fmt.Println("------------------- Sucessfully connected to the database! -------------------")

	return db, nil
}
