package util

import (
	"elkFormatter/constant"
	"encoding/json"
	"strings"
	"time"
)

func ParseJsonBody(bodyJson string, hasError bool) string {
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(bodyJson), &result)
	if err != nil {
		return bodyJson
	}

	delete(result, "rq_header")

	indentBody, _ := json.MarshalIndent(result, "", "    ")
	lines := strings.Split(string(indentBody), "\n")

	if hasError {
		for i := range lines {
			lines[i] = constant.Red + lines[i] + constant.Reset
		}
	}

	return strings.Join(lines, "\n")
}

func FormatTimeToString(data time.Time) string {
	date := data.Format("2006-01-02")
	timeStamp := data.Format("2006-01-02 15:04:05.000")
	return date + " " + timeStamp
}
