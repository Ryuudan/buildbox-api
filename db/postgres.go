package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Pyakz/buildbox-api/ent/generated"
	_ "github.com/lib/pq" // Import the pq driver
)

func PostgresConnect() (*generated.Client, error) {

	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	client, err := generated.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname,
	))

	if err != nil {
		return nil, err
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Successfully connected to the database
	log.Printf("Sucessfully connected to the database!")

	return client, nil
}
