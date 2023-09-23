package database

import (
	"context"
	"log"
	"os"

	"github.com/Pyakz/buildbox-api/ent/generated"
	_ "github.com/lib/pq" // Import the pq driver
)

func PostgresConnect() (*generated.Client, error) {

	client, err := generated.Open("postgres", os.Getenv("POSTGRES_CONNECTION_STRING"))

	if err != nil {
		return nil, err
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Successfully connected to the database
	log.Println("✅ Sucessfully connected to the Postgres Database!")

	return client, nil
}
