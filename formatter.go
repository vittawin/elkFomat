package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func setInput() {
	file, err := getInput()
	if err != nil {
		panic(err)
	}

	for file.Scan() {
		curRow := file.Text()
		curLog := LogStruct{}
		if err := json.Unmarshal([]byte(curRow), &curLog); err != nil {
			fmt.Println("cannot unmarshal :", err)
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
		description: parsedRow.ServiceName + " : " + parsedRow.EntryModule,
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
		return ""
	}

	delete(result, "rq_header")

	indentBody, _ := json.MarshalIndent(result, "", "    ")
	return string(indentBody)
}

func parseLogBody(row string) (string, error) {
	log := LogStruct{}
	if err := json.Unmarshal([]byte(row), &log); err != nil {
		fmt.Println("cannot unmarshal :", err)
	}

	return log.Body, nil
}
