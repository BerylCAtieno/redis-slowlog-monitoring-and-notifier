package database

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/joho/godotenv"
)

// Global Redis client
var RedisClient *redis.Client
var ctx = context.Background()

// Initialize Redis connection
func InitRedis() {
	// Load .env file (only for local development)
	_ = godotenv.Load()

	// Get Redis URL from environment variables
	redisURL := os.Getenv("UPSTASH_REDIS_URL")
	if redisURL == "" {
		log.Fatal("UPSTASH_REDIS_URL is not set in the environment")
	}

	// Parse the Redis URL
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("Failed to parse Redis URL: %v", err)
	}

	// Create a new Redis client
	RedisClient = redis.NewClient(opt)
	log.Println("✅ Connected to Redis successfully")
}

// Close Redis connection (to be called when the app shuts down)
func CloseRedis() {
	if RedisClient != nil {
		RedisClient.Close()
		log.Println("🛑 Redis connection closed")
	}
}

// Helper function to get Redis context
func GetRedisContext() context.Context {
	return ctx
}
