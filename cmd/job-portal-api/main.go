package main

import (
	"fmt"
	"net/http"
	"os"
	"project/internal/auth"
	"project/internal/database"
	"project/internal/handlers"
	"project/internal/repository"
	"project/internal/services"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

// ===================main()===============================
func main() {

	// starting the application.
	err := startApp()

	// Check if there was an error while starting the application.
	if err != nil {
		// Log a panic and include the error details if an error occurred.
		log.Panic().Err(err).Send()
	}

}

// =============================startApp()============================
func startApp() error {

	// Log that the application has started.
	log.Info().Msg("Starting job portal application")

	// Load the private RSA key from a file.
	privatePEM, err := os.ReadFile(`C:\Users\ORR Training 3\Desktop\job-portal-api\private.pem`)
	if err != nil {
		return fmt.Errorf("failed to load private.pem file: %w", err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return fmt.Errorf("failed to parse private key from bytes: %w", err)
	}

	// Load the public RSA key from a file.
	publicPEM, err := os.ReadFile(`C:\Users\ORR Training 3\Desktop\job-portal-api\pubkey.pem`)
	if err != nil {
		return fmt.Errorf("failed to load pubkey.pem file: %w", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPEM)
	if err != nil {
		return fmt.Errorf("failed to parse public key from bytes: %w", err)
	}

	// Create an authentication instance using the loaded keys.
	authInstance, err := auth.NewAuth(privateKey, publicKey)
	if err != nil {
		return fmt.Errorf("failed to create authentication instance: %w", err)
	}

	// Establish a connection to the database.
	db, err := database.Connection()
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Initialize a repository with the database connection.
	repo, err := repository.NewRepo(db)
	if err != nil {
		return fmt.Errorf("failed to initialize repository: %w", err)
	}

	// Create a service instance using the repository.
	serviceInstance, err := services.NewService(repo, repo)
	if err != nil {
		return fmt.Errorf("failed to create service instance: %w", err)
	}

	// Configure and start the HTTP server.
	api := http.Server{
		Addr:    ":8085",
		Handler: handlers.API(authInstance, serviceInstance),
	}
	api.ListenAndServe()

	return nil

}
