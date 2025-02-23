package database

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestNewRedisClient(t *testing.T) {
	// Load .env file for testing
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	// Check if REDIS_URL is set
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		t.Fatal("REDIS_URL environment variable not set")
	}

	// Test client creation
	client, err := NewRedisClient()
	if err != nil {
		t.Fatalf("Failed to create Redis client: %v", err)
	}
	defer client.Close()

	// Test connection with ping
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Ping(ctx)
	if err != nil {
		t.Fatalf("Failed to ping Redis: %v", err)
	}
}

func TestRedisClientWithInvalidURL(t *testing.T) {
	// Temporarily set invalid Redis URL
	os.Setenv("REDIS_URL", "invalid://localhost:6379")
	defer os.Setenv("REDIS_URL", "") // Reset after test

	// Test client creation with invalid URL
	_, err := NewRedisClient()
	if err == nil {
		t.Error("Expected error with invalid Redis URL, got nil")
	}
}