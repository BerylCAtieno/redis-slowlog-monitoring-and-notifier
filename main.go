package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/BerylCAtieno/redis-slowlog-monitor/monitor" // Import your package
)

func main() {
	ctx := context.Background()

	// Connect to the local Redis server
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Default Redis address
	})

	// Test fetching slow logs
	slowLogs := monitor.FetchSlowLogs(ctx, client)
	if slowLogs == nil {
		log.Println("No slow logs found or an error occurred.")
		return
	}

	// Print slow logs
	fmt.Println("Fetched slow logs:")
	for _, logEntry := range slowLogs {
		fmt.Printf("ID: %v | Timestamp: %v | Duration: %v μs | Command: %v\n",
			logEntry[0], logEntry[1], logEntry[2], logEntry[3])
	}
}
