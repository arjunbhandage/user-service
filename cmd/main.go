package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user-service/database"
	"user-service/user"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection

	url := GetEnv("DB_URL", "postgres://admin:password123@localhost:5432/users?sslmode=disable")
	db := database.InitDB(url)

	// Run migrations
	database.RunMigrations(url)

	// Initialize routes
	r := gin.Default()
	user.RegisterRoutes(r, user.NewUserService(database.NewUserRepo(db)))

	// Start the server in a separate goroutine
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		log.Println("Starting server on :8080")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Create a channel to listen for OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Wait for a signal
	<-stop
	log.Println("Shutting down server...")

	// Create a context with a timeout to allow for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exited gracefully")
}

// GetEnv returns the value of an environment variable or a default value if not set
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
