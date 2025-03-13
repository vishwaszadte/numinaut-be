package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/vishwaszadte/numinaut-be/internal/handler"
	"github.com/vishwaszadte/numinaut-be/internal/middleware"
	"github.com/vishwaszadte/numinaut-be/internal/repository"
	"github.com/vishwaszadte/numinaut-be/internal/service"
	"github.com/vishwaszadte/numinaut-be/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	isLogLevelDebug := os.Getenv("LOG_LEVEL") == "debug"

	logger.InitializeZapLogger(isLogLevelDebug)
	defer logger.Sync()

	// Initialize database connection
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		logger.Fatal("Unable to create connection pool", zap.Error(err))
	}
	defer dbpool.Close()

	// Initialize repository
	queries := repository.New(dbpool)

	// Initialize service
	expressionService := service.NewExpressionService(queries)

	// Initialize handler
	expressionHandler := handler.NewExpressionHandler(expressionService)

	// Initialize router
	router := mux.NewRouter()

	// Add middleware
	router.Use(middleware.LoggingMiddleware)

	// Register routes
	expressionHandler.RegisterRoutes(router)

	// Create server
	server := &http.Server{
		Addr:         "localhost:8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		logger.Info("Server is starting on port 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Could not listen on port 8080", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until the timeout deadline
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Error during server shutdown", zap.Error(err))
	}

	logger.Info("Server gracefully stopped")
}
