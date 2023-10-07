package database

import (
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

	// Successfully connected to the database
	log.Println("✅ Sucessfully connected to the Postgres Database!")

	return client, nil
}
