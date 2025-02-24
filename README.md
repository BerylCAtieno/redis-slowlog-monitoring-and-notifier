# Redis Slowlog Monitor Integration for Telex

A real-time Redis slow query monitoring integration for Telex messaging platform. This application monitors Redis slow logs and sends notifications to your Telex channels when queries exceed specified thresholds.

## Table of Content

- [Overview](#overview)
- [Features](#features)
- [About Telex Integration](#about-telex-integration)
- [Directory Structure](#directory-structure)
- [API Documentation](#api-documentation)
- [Installation and Usage](#installation-and-usage)
- [Manual Testing](#manual-testing)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)
- [Author](#author)


## Overview
Redis SlowLog Monitor is a Go application that integrates with Telex to provide real-time monitoring of Redis slow queries. It fetches slow log entries from your Redis instance, analyzes them based on configurable thresholds, and sends formatted alerts to your specified Telex channels.

## Features
- Monitors Redis slow queries based on a configurable threshold
- Sends alerts based on user-defined settings
- Supports cron-based query check intervals
- Customizable sensitivity levels and notification preferences

## About Telex Integration
Telex is a messaging platform that supports custom integrations. This application implements a "modifier" type integration, which means it:

- Receives messages via webhooks
- Processes them according to configured settings
- Returns formatted responses to Telex channels

### Integration Settings

The application supports custom settings:

| Setting                  | Type        | Description                                      | Default Value |
|--------------------------|------------|--------------------------------------------------|---------------|
| Enable Notifications     | Checkbox   | Enable or disable alerts                        | `yes`         |
| Slow Query Threshold (ms)| Number     | Redis slow query threshold in milliseconds      | `100`         |

## Directory Structure

```
redis-slowlog-monitor/
├── api/
│   └── handlers.go          # HTTP handlers for the integration
├── database/
│   └── redis.go            # Redis client configuration and connection
|   └── redis_test.go        # Redis client tests
├── models/
│   └── models.go           # Data structures and types
├── monitor/
│   ├── monitor.go            # Slow log fetching logic
│   └── alert.go           # Alert message formatting
├── .gitignore                  
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
├── main.go                 # Application entry point
└── README.md              # This documentation
```

## API Documentation

### Integration Configuration Endpoint

```
GET /integration.json
```
Returns the integration configuration for Telex.
Response example:

```json
{
  "data": {
    "author": "Beryl Atieno",
    "descriptions": {
      "app_name": "Redis SlowLog Monitor",
      "app_description": "A monitoring tool that detects Redis slow queries..."
    },
    "integration_type": "modifier",
    "settings": [
      {
        "label": "Enable Notifications",
        "type": "checkbox",
        "required": true,
        "default": "Yes"
      },
      {
        "label": "Slow Query Threshold (ms)",
        "type": "number",
        "required": true,
        "default": "100"
      }
    ]
  }
}
```
### Telex Alert Formatting Endpoint

```
POST /format-alert
```
Processes Redis slow logs and returns formatted alerts.
Request body:

```json
{
  "channel_id": "channel-123",
  "settings": [
    {
      "label": "Enable Notifications",
      "type": "checkbox",
      "default": "Yes",
      "required": true
    },
    {
      "label": "Slow Query Threshold (ms)",
      "type": "number",
      "default": 100,
      "required": true
    }
  ],
  "message": "Test message"
}
```
Response example:

```json
{
  "event_name": "message_formatted",
  "message": "⚠️ Redis Slow Query Alert ⚠️\n\nTime: 2025-02-24 15:04:05\nDuration: 150ms\nCommand: SET key value\n-------------------",
  "status": "success",
  "username": "Redis SlowLog Monitor"
}
```
## Installation and Usage

### Prerequisites
- Go (1.18 or later)
- Redis server instance
- Telex account with integration configuration access

### Set-Up Instructions

1. Clone the repository:

```sh
git clone https://github.com/BerylCAtieno/redis-slowlog-monitor.git
cd redis-slowlog-monitor
```
2. Install dependencies

```sh
go mod download
```

3. Create a .env file

```env
REDIS_URL=redis://username:password@host:port
PORT=8080
BASE_URL=https://your-application-url.com
```
4. Build the Application

```sh
go build -o redis-monitor
```

5. Start the server
```sh
./redis-monitor
```

## Manual Testing

Test the API endpoints using curl:

1. Test Integration Configuration 

```sh
curl http://redis-monitoring-telex-integration.onrender.com/format-alert/integration.json
```

2. Test Alert Formatting

```sh
curl -X POST "http://redis-monitoring-telex-integration.onrender.com/format-alert" \
     -H "Content-Type: application/json" \
     -d '{
       "channel_id": "019532f7-9e79-7574-89f5-ab7c3e6fa9c2",
       "settings": [
         {
           "label": "Enable Notifications",
           "type": "checkbox",
           "default": "Yes",
           "required": true
         },
         {
           "label": "Slow Query Threshold (ms)",
           "type": "number",
           "default": 100,
           "required": true
         }
       ],
       "message": "Test message"
     }'
```
## Contributing

### Getting Started

1. Fork the repository
2. Create a feature branch:

```sh
git checkout -b feature/your-feature-name
```

3. Make your changes
4. Run tests to ensure everything works
5. Commit your changes

```sh
git commit -m "Add some feature"
```

6. Push to your fork:

```sh
git push origin feature/your-feature-name
```

7. Create a Pull Request

### Development Guidelines

- Follow Go best practices and coding standards
- Write unit tests for new features
- Update documentation as needed
- Keep commits atomic and well-described
- Use descriptive branch names


## License
This project is licensed under the MIT License.

## Acknowledgments

- Redis team for the excellent slow log feature
- Telex team for the integration platform

## Author
[BerylCAtieno](https://github.com/BerylCAtieno)
