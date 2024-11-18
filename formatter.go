package main

import (
	"bufio"
	"elkFormatter/constant"
	"elkFormatter/logUtil"
	"elkFormatter/util"
	"encoding/json"
	"github.com/charmbracelet/bubbles/list"
	"os"
)

var data []string
var rows = make(map[string][]logUtil.LogStruct) //keys is jobId of log
var keys = make([]string, 0)
var dataLogList []list.Item

func setInput() {
	file, err := getInput()
	if err != nil {
		panic(err)
	}

	for file.Scan() {
		curRow := file.Text()
		curLog := logUtil.LogStruct{}
		if err := json.Unmarshal([]byte(curRow), &curLog); err != nil {
			//fmt.Println("cannot unmarshal :", err, "#####")
			continue
		}

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

func addRows(logRow logUtil.LogStruct, rawRow string) {
	parsedRow := logUtil.ParseRowData(logRow)

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

func parseDataToString(row logUtil.LogStruct) string {
	switch row.EntryModule {
	case constant.Kafka:
		return logUtil.KafkaLog(row)
	case constant.Rest:
		return logUtil.RestLog(row)
	case constant.Event:
		return logUtil.EventLog(row)
	default:
		return logUtil.DefaultLog(row)
	}
}

func parseLogBody(row string) (string, error) {
	log := logUtil.LogStruct{}
	if err := json.Unmarshal([]byte(row), &log); err != nil {
	}

	result := parseDataToString(log) + "\n" +
		"body : " + util.ParseJsonBody(log.Body) + "\n"
	return result, nil
}
