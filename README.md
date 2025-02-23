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
## Integration json
```
https://redis-monitoring-telex-integration.onrender.com/integration.json
```

## Configuration

The application supports custom settings:

| Setting                  | Type        | Description                                      | Default Value |
|--------------------------|------------|--------------------------------------------------|---------------|
| Enable Notifications     | Checkbox   | Enable or disable alerts                        | `yes`         |
| Slow Query Threshold (ms)| Number     | Redis slow query threshold in milliseconds      | `100`         |


## Test /format-message

```
curl -X POST "http://https://redis-monitoring-telex-integration.onrender.com/format-alert" \
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

## License
This project is licensed under the MIT License.

## Author
[BerylCAtieno](https://github.com/BerylCAtieno)
