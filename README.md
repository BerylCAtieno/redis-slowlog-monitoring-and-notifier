# Redis Slowlog Monitor

Redis Slowlog Monitor is a Go-based application that tracks slow queries in Redis and provides configurable alerts for monitoring performance.

## Features
- Monitors Redis slow queries based on a configurable threshold
- Sends alerts based on user-defined settings
- Supports cron-based query check intervals
- Customizable sensitivity levels and notification preferences

## Installation

### Prerequisites
- Go (1.18 or later)
- Redis server

### Clone the Repository
```sh
git clone https://github.com/BerylCAtieno/redis-slowlog-monitor.git
cd redis-slowlog-monitor
```

### Build the Application
```sh
go build -o slowlog-monitor .
```

### Run the Application
```sh
./slowlog-monitor
```

## Configuration

The application supports custom settings:

| Setting                  | Type        | Description                                      | Default Value |
|--------------------------|------------|--------------------------------------------------|---------------|
| Query Check Interval     | Text       | Cron expression for monitoring queries          | `0 * * * *`   |
| Enable Notifications     | Checkbox   | Enable or disable alerts                        | `yes`         |
| Slow Query Threshold (ms)| Number     | Redis slow query threshold in milliseconds      | `100`         |
| Sensitivity Level        | Dropdown   | Sensitivity level (High, Medium, Low)           | `Medium`      |
| Alert Recipients        | Multi-checkbox | Recipients for alerts                         | `Admin`       |

## Example Usage

```go
msg := models.Message{
    ChannelID: "01952e73-0caa-7bf8-b2ce-aca0f1a02acd",
    Settings: []models.Setting{
        {Label: "Query Check Interval", Type: "text", Default: "0 * * * *", Required: true},
        {Label: "Enable Notifications", Type: "checkbox", Default: "yes", Required: true},
        {Label: "Slow Query Threshold (ms)", Type: "number", Default: 100, Required: true},
        {Label: "Sensitivity Level", Type: "dropdown", Default: "Medium", Required: false},
        {Label: "Alert Recipients", Type: "multi-checkbox", Default: "Admin", Required: false},
    },
    Message: "SlowLog detected!!",
}
settingsProcessing(msg)
```

## License
This project is licensed under the MIT License.

## Author
[BerylCAtieno](https://github.com/BerylCAtieno)
