package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Pyakz/buildbox-api/database"
	"github.com/Pyakz/buildbox-api/internal/routers"
	"github.com/Pyakz/buildbox-api/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func init() {
	// Load environment variables here
	if err := utils.LoadEnvironmentVariables(); err != nil {
		fmt.Printf("Failed to load environment variables: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	// init redis
	redis_client := database.RedisClient()
	db_client, err := database.PostgresConnect()

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	defer db_client.Close()

	// Set up router
	app := chi.NewRouter()
	app.Use(middleware.Heartbeat("/ping"))
	app.Use(middleware.RequestID)
	app.Use(middleware.RealIP)
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)
	app.Use(middleware.Throttle(100))

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Api-Version"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	app.Use(cors.Handler)

	app.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte("This is Build Box API"))
		if err != nil {
			log.Println("Error writing response:", err)
		}
	})

	app.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	routers.PrivateInitializeRoutes(db_client, app)
	routers.PublicInitializeRoutes(db_client, redis_client, app)

	// Start server
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: app,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("❌ Error starting server: %v", err)
			os.Exit(1)
		}
	}()

	log.Printf("✅ Server started on port %s", os.Getenv("PORT"))

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("❌ Error shutting down server: %v", err)
	}
}
