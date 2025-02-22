package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/BerylCAtieno/redis-slowlog-monitor/models"
)


func SendNotification(slowLogs []models.SlowLogEntry) {
	if len(slowLogs) == 0 {
		log.Println("No slow queries detected. No notification sent.")
		return
	}

	// Format the message
	message := "Slow queries detected:\n"
	for _, log := range slowLogs {
		message += fmt.Sprintf("ID: %s, Command: %s, Duration: %dµs\n", log.ID, log.Command, log.Duration)
	}

	// Construct the API payload
	payload := models.APIRequestPayload{
		ChannelID: "01952e73-0caa-7bf8-b2ce-aca0f1a02acd",
		Settings: []struct {
			Label       string      `json:"label"`
			Type        string      `json:"type"`
			Description string      `json:"description"`
			Default     interface{} `json:"default"`
			Required    bool        `json:"required"`
		}{
			{
				Label:       "maxMessageLength",
				Type:        "number",
				Description: "Set the maximum length for incoming messages to format.",
				Default:     200,
				Required:    true,
			},
		},
		Message: message,
	}

	// Convert to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return
	}

	// Send HTTP POST request
	resp, err := http.Post("http://your-api.com/format-message", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Failed to send notification:", err)
		return
	}
	defer resp.Body.Close()

	log.Println("Notification sent. Status:", resp.Status)
}
