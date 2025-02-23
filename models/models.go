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

// Integration settings struct

type Settings struct {
	Label    string      `json:"label"`
	Type     string      `json:"type"`
	Default  interface{} `json:"default"`
	Required bool        `json:"required"`
}

// Response payload
type Response struct {
	EventName string `json:"event_name"`
	Message   string `json:"message"`
	Status    string `json:"status"`
	Username  string `json:"username"`
}


// alert message struct
type Message struct {
	ChannelID string    `json:"channel_id"`
	Settings  []Settings `json:"settings"`
	Message   string    `json:"message"`
}

// Integration config json model

type IntegrationConfig struct {
    Data struct {
        Author     string `json:"author"`
        Date       struct {
            CreatedAt string `json:"created_at"`
            UpdatedAt string `json:"updated_at"`
        } `json:"date"`
        Descriptions struct {
            AppDescription  string `json:"app_description"`
            AppLogo        string `json:"app_logo"`
            AppName        string `json:"app_name"`
            AppURL         string `json:"app_url"`
            BackgroundColor string `json:"background_color"`
        } `json:"descriptions"`
        IntegrationCategory string `json:"integration_category"`
        IntegrationType    string `json:"integration_type"`
        IsActive          bool   `json:"is_active"`
        Output            []struct {
            Label string `json:"label"`
            Value bool   `json:"value"`
        } `json:"output"`
        KeyFeatures []string `json:"key_features"`
        Permissions struct {
            MonitoringUser struct {
                AlwaysOnline bool   `json:"always_online"`
                DisplayName string `json:"display_name"`
            } `json:"monitoring_user"`
        } `json:"permissions"`
        Settings    []Settings `json:"settings"`
        TargetURL   string    `json:"target_url"`
        TickURL     string    `json:"tick_url"`
        Website     string    `json:"website"`
    } `json:"data"`
}