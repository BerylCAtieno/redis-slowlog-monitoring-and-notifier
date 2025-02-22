package models

// SlowLogEntry defines the structure of a slow log
type SlowLogEntry struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Duration  int64  `json:"duration"`
	Command   string `json:"command"`
}

// APIRequestPayload defines the structure of the /format-message POST request
type APIRequestPayload struct {
	ChannelID string `json:"channel_id"`
	Settings  []struct {
		Label       string      `json:"label"`
		Type        string      `json:"type"`
		Description string      `json:"description"`
		Default     interface{} `json:"default"`
		Required    bool        `json:"required"`
	} `json:"settings"`
	Message string `json:"message"`
}