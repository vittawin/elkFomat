package util

import (
	"encoding/json"
	"time"
)

func ParseJsonBody(bodyJson string) string {
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(bodyJson), &result)
	if err != nil {
		return bodyJson
	}

	delete(result, "rq_header")

	indentBody, _ := json.MarshalIndent(result, "", "    ")
	return string(indentBody)
}

func FormatTimeToString(data time.Time) string {
	date := data.Format("2006-01-02")
	timeStamp := data.Format("2006-01-02 15:04:05.000")
	return date + " " + timeStamp
}
