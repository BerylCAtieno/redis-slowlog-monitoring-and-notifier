package monitor

import (
	"fmt"
	"time"
)

func FormatSlowLogAlert(rows [][]interface{}, threshold int) string {
	if len(rows) == 0 {
		return "No slow queries detected."
	}

	alert := "⚠️ Redis Slow Query Alert ⚠️\n\n"
	for _, row := range rows {
		duration := row[2].(int64)
		if duration >= int64(threshold)*1000 { // Convert ms to microseconds
			timestamp := time.Unix(row[1].(int64), 0).Format("2006-01-02 15:04:05")
			command := row[3].(string)
			
			alert += fmt.Sprintf("Time: %s\n", timestamp)
			alert += fmt.Sprintf("Duration: %dms\n", duration/1000)
			alert += fmt.Sprintf("Command: %s\n", command)
			alert += "-------------------\n"
		}
	}
	
	return alert
}