package monitor

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func FetchSlowLogs(ctx context.Context, client *redis.Client) [][]interface{} {
	slowLogs, err := client.Do(ctx, "SLOWLOG", "GET", 10).Result()
	if err != nil {
		log.Println("Failed to fetch slow logs:", err)
		return nil
	}

	var rows [][]interface{}
	if logs, ok := slowLogs.([]interface{}); ok {
		for _, logEntry := range logs {
			if logDetails, ok := logEntry.([]interface{}); ok && len(logDetails) >= 4 {
				id := logDetails[0]          // Slow log ID
				timestamp := logDetails[1]   // Unix timestamp
				duration := logDetails[2]    // Execution time (microseconds)
				command := logDetails[3]     // Command run

				row := []interface{}{id, timestamp, duration, fmt.Sprintf("%v", command)}
				rows = append(rows, row)
			}
		}
	}
	return rows
}