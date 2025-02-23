package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/BerylCAtieno/redis-slowlog-monitor/database"
	"github.com/BerylCAtieno/redis-slowlog-monitor/models"
	"github.com/BerylCAtieno/redis-slowlog-monitor/monitor"
)

func HandleIncomingMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var msgReq models.Message
	if err := json.NewDecoder(r.Body).Decode(&msgReq); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	log.Printf("Received message request with settings: %+v", msgReq.Settings)

	// Extract settings
	threshold := 100 // Default threshold
	notificationsEnabled := true

	for _, setting := range msgReq.Settings {
        log.Printf("Processing setting: %s with value: %v", setting.Label, setting.Default)
        switch setting.Label {
        case "Slow Query Threshold (ms)":
            if val, ok := setting.Default.(float64); ok {
                threshold = int(val)
                log.Printf("Set threshold to: %d", threshold)
            }
        case "Enable Notifications":
            if val, ok := setting.Default.(string); ok {
                notificationsEnabled = val == "Yes"
                log.Printf("Notifications enabled: %v", notificationsEnabled)
            }
        }
    }

	// Initialize Redis client using our database package
	redisClient, err := database.NewRedisClient()
	if err != nil {
		log.Printf("Failed to connect to Redis: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer redisClient.Close()

	// Fetch slow logs
	ctx := context.Background()
	slowLogs := monitor.FetchSlowLogs(ctx, redisClient.Client)

	// Format alert message
	formattedMessage := monitor.FormatSlowLogAlert(slowLogs, threshold)
	log.Printf("Formatted message before notification check: %s", formattedMessage)

	// Only send if notifications are enabled and there are slow queries
	if !notificationsEnabled {
		log.Printf("Notifications disabled, returning empty message")
		formattedMessage = ""
	} else if formattedMessage == "No slow queries detected." {
		log.Printf("No slow queries detected")
		// You might want to still return this message instead of empty string
		// formattedMessage = "No slow queries detected."
	}

	response := models.Response{
		EventName: "message_formatted",
		Message:   formattedMessage,
		Status:    "success",
		Username:  "Redis SlowLog Monitor",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handler for integration config

func GetIntegrationConfig() models.IntegrationConfig {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "https://redis-monitoring-telex-integration.onrender.com"
	}

	return models.IntegrationConfig{
		Data: struct {
			Author string `json:"author"`
			Date   struct {
				CreatedAt string `json:"created_at"`
				UpdatedAt string `json:"updated_at"`
			} `json:"date"`
			Descriptions struct {
				AppDescription  string `json:"app_description"`
				AppLogo         string `json:"app_logo"`
				AppName         string `json:"app_name"`
				AppURL          string `json:"app_url"`
				BackgroundColor string `json:"background_color"`
			} `json:"descriptions"`
			IntegrationCategory string   `json:"integration_category"`
			IntegrationType     string   `json:"integration_type"`
			IsActive            bool     `json:"is_active"`
			KeyFeatures         []string `json:"key_features"`
			Permissions         struct {
				MonitoringUser struct {
					AlwaysOnline bool   `json:"always_online"`
					DisplayName  string `json:"display_name"`
				} `json:"monitoring_user"`
			} `json:"permissions"`
			Settings  []models.Settings `json:"settings"`
			TargetURL string            `json:"target_url"`
			TickURL   string            `json:"tick_url"`
			Website   string            `json:"website"`
		}{
			Author: "Beryl Atieno",
			Date: struct {
				CreatedAt string `json:"created_at"`
				UpdatedAt string `json:"updated_at"`
			}{
				CreatedAt: "2025-02-22",
				UpdatedAt: "2025-02-22",
			},
			Descriptions: struct {
				AppDescription  string `json:"app_description"`
				AppLogo         string `json:"app_logo"`
				AppName         string `json:"app_name"`
				AppURL          string `json:"app_url"`
				BackgroundColor string `json:"background_color"`
			}{
				AppDescription:  "A monitoring tool that detects Redis slow queries and notifies administrators in real time.",
				AppLogo:         "https://i.imgur.com/7JQ7JEX.png",
				AppName:         "Redis SlowLog Monitor",
				AppURL:          "https://redis-monitoring-telex-integration.onrender.com/format-alert",
				BackgroundColor: "#ffffff",
			},
			IntegrationCategory: "Performance Monitoring",
			IntegrationType:     "modifier",
			IsActive:            true,
			KeyFeatures: []string{
				"Real-time detection of Redis slow queries.",
				"Customizable alert settings for different sensitivity levels.",
				"Integration with Telex message channels.",
				"Detailed information for each slowlog.",
			},
			Permissions: struct {
				MonitoringUser struct {
					AlwaysOnline bool   `json:"always_online"`
					DisplayName  string `json:"display_name"`
				} `json:"monitoring_user"`
			}{
				MonitoringUser: struct {
					AlwaysOnline bool   `json:"always_online"`
					DisplayName  string `json:"display_name"`
				}{
					AlwaysOnline: true,
					DisplayName:  "Redis SlowLog Monitoring Bot",
				},
			},
			Settings: []models.Settings{
				{
					Label:    "Enable Notifications",
					Type:     "checkbox",
					Required: true,
					Default:  "Yes",
				},
				{
					Label:    "Slow Query Threshold (ms)",
					Type:     "number",
					Required: true,
					Default:  float64(100),
				},
			},
			TargetURL: baseURL + "/format-alert",
			TickURL:   baseURL + "/format-alert",
			Website:   "https://redis-monitoring-telex-integration.onrender.com/",
		},
	}
}

func HandleIntegrationConfig(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	config := GetIntegrationConfig()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config)
}
