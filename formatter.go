package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	"os"
	"time"
)

const (
	rqHeader = "rq_header"
	kafka    = "KAFKA"
	rest     = "REST"
	event    = "EVENT"

	reset         = "\033[0m"
	black         = "\033[30m"
	red           = "\033[31m"
	green         = "\033[32m"
	yellow        = "\033[33m"
	blue          = "\033[34m"
	magenta       = "\033[35m"
	cyan          = "\033[36m"
	white         = "\033[37m"
	brightBlack   = "\033[90m"
	brightRed     = "\033[91m"
	brightGreen   = "\033[92m"
	brightYellow  = "\033[93m"
	brightBlue    = "\033[94m"
	brightMagenta = "\033[95m"
	brightCyan    = "\033[96m"
	brightWhite   = "\033[97m"
)

var data []string
var rows = make(map[string][]LogStruct) //keys is jobId of log
var keys = make([]string, 0)
var dataLogList []list.Item

func setInput() {
	file, err := getInput()
	if err != nil {
		panic(err)
	}

	for file.Scan() {
		curRow := file.Text()
		curLog := LogStruct{}
		if err := json.Unmarshal([]byte(curRow), &curLog); err != nil {
			fmt.Println("cannot unmarshal :", err, "#####")
			continue
		}

		curLog.ErrorMessage, _ = formatErrorMessage(curLog.ErrorMessage)

		if curLog.Module == "biz" {
			continue
		}
		addRows(curLog, curRow)

	}
}

func getInput() (*bufio.Scanner, error) {
	file := os.Stdin
	return bufio.NewScanner(file), nil
}

func addRows(logRow LogStruct, rawRow string) {
	parsedRow := parseRowData(logRow)

	if _, ok := rows[parsedRow.JobID]; !ok {
		keys = append(keys, parsedRow.JobID)
	}

	dataLogList = append(dataLogList, Log{
		jobId:       parsedRow.JobID,
		data:        rawRow,
		title:       parsedRow.ServiceName,
		description: parseDataToString(parsedRow),
	})

	rows[parsedRow.JobID] = append(rows[parsedRow.JobID], parsedRow)
}

func parseRowData(row LogStruct) LogStruct {
	row.Body = parseJsonBody(row.Body)

	if row.JobID == "" {
		//row.JobID = findJobId(row.Body)
	}

	return row
}

func parseJsonBody(bodyJson string) string {
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(bodyJson), &result)
	if err != nil {
		return bodyJson
	}

	delete(result, "rq_header")

	indentBody, _ := json.MarshalIndent(result, "", "    ")
	return string(indentBody)
}

func parseLogBody(row string) (string, error) {
	log := LogStruct{}
	if err := json.Unmarshal([]byte(row), &log); err != nil {
		fmt.Println("cannot unmarshal :", err, "#####", row)
	}

	result := parseDataToString(log) + "\n" +
		"body : " + parseJsonBody(log.Body) + "\n"
	return result, nil
}

func parseDataToString(row LogStruct) string {
	switch row.EntryModule {
	case kafka:
		return kafkaLog(row)
	case rest:
		return restLog(row)
	case event:
		return eventLog(row)
	default:
		return defaultLog(row)
	}
}

func formatTimeToString(data time.Time) string {
	date := data.Format("2006-01-02")
	timeStamp := data.Format("2006-01-02 15:04:05.000")
	return date + " " + timeStamp
}

func restFormat(row LogStruct) string {
	if row.Module == "client" {
		return "ToPath : " + row.ToPath + "\n"
	}
	return ""
}

func restPath(row LogStruct) string {
	if row.OriginalPath != "" {
		return row.OriginalPath
	}
	return row.Path
}

func logErr(row LogStruct) string {
	if row.ErrorMessage != nil {
		errMassage, ok := row.ErrorMessage.([]ErrorStruct)
		if ok && len(errMassage) != 0 {
			fmt.Println("Error : ", errMassage)
			return "Error : " + errMassage[0].ErrorDetail
		}
	}
	return ""
}

func kafkaLog(row LogStruct) string {
	errorMessage, ok := row.ErrorMessage.([]ErrorStruct)
	if ok && errorMessage != nil {
		fmt.Println("Error : ", row.ErrorMessage)
		return red + "Time : " + formatTimeToString(row.Timestamp) + "\n" +
			red + "Module : " + row.Module + " | " + row.Type + "\n" +
			red + row.ServiceName + " | " + row.EntryModule + "\n" +
			red + row.Topic + " : " + row.Path + "\n" +
			red + logErr(row) + reset
	} else {
		return green + "Time : " + formatTimeToString(row.Timestamp) + "\n" +
			green + "Module : " + row.Module + " | " + row.Type + "\n" +
			green + row.ServiceName + " | " + row.EntryModule + "\n" +
			green + row.Topic + " : " + row.Path + "\n" +
			green + logErr(row)
	}
}

func restLog(row LogStruct) string {
	if row.OriginalPath == "" {
		row.OriginalPath = row.EventName
	}

	errorMessage, ok := row.ErrorMessage.([]ErrorStruct)
	if ok && errorMessage != nil {
		fmt.Println("Error : ", row.ErrorMessage)
		return red + "Time : " + formatTimeToString(row.Timestamp) + "\n" +
			red + "Module : " + row.Module + " | " + row.Type + "\n" +
			red + restFormat(row) +
			red + row.ServiceName + " | " + row.EntryModule + "\n" +
			red + row.Method + " : " + restPath(row) + "\n" +
			red + logErr(row) + reset
	} else {
		return blue + "Time : " + formatTimeToString(row.Timestamp) + "\n" +
			blue + "Module : " + row.Module + " | " + row.Type + "\n" +
			blue + restFormat(row) +
			blue + row.ServiceName + " | " + row.EntryModule + "\n" +
			blue + row.Method + " : " + restPath(row) + "\n" +
			blue + logErr(row) + reset
	}
}

func eventLog(row LogStruct) string {
	errorMessage, ok := row.ErrorMessage.([]ErrorStruct)
	if ok && errorMessage != nil {
		return red + "Time : " + formatTimeToString(row.Timestamp) + "\n" +
			red + "Module : " + row.Module + " | " + row.Type + "\n" +
			red + row.ServiceName + " | " + row.EntryModule + "\n" +
			red + row.Path + " : " + row.ToPath + "\n" +
			red + logErr(row) + reset
	} else {
		return yellow + "Time : " + formatTimeToString(row.Timestamp) + "\n" +
			yellow + "Module : " + row.Module + " | " + row.Type + "\n" +
			yellow + row.ServiceName + " | " + row.EntryModule + "\n" +
			yellow + row.Path + " : " + row.ToPath + "\n" +
			yellow + logErr(row)
	}
}

func defaultLog(row LogStruct) string {
	errorMessage, ok := row.ErrorMessage.([]ErrorStruct)
	if ok && errorMessage != nil {
		return red + "Time : " + formatTimeToString(row.Timestamp) + "\n" +
			red + "Module : " + row.Module + " | " + row.Type + "\n" +
			red + row.ServiceName + " | " + row.EntryModule + "\n" +
			red + row.Topic + " : " + row.Path + "\n" +
			red + logErr(row) + reset
	} else {
		return "Time : " + formatTimeToString(row.Timestamp) + "\n" +
			"Module : " + row.Module + " | " + row.Type + "\n" +
			row.ServiceName + " | " + row.EntryModule + "\n" +
			row.Topic + " : " + row.Path + "\n" +
			logErr(row)
	}
}

func formatErrorMessage(input any) ([]ErrorStruct, error) {
	var result []ErrorStruct
	if input != nil {
		inputStr, ok := input.(string)
		if inputStr != "" && ok {
			err := json.Unmarshal([]byte(inputStr), &result)
			if err != nil {
				return nil, err
			}

			fmt.Println("Error : ", result)
			return result, nil
		}
		if !ok {
			fmt.Println("Errorsssssss : ", input)
		}
	}

	return nil, nil
}
