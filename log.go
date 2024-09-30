package main

import "time"

type LogStruct struct {
	Timestamp    time.Time `json:"timestamp"`
	Level        string    `json:"level"` //type log?
	DomainCode   string    `json:"domainCode"`
	DomainName   string    `json:"domainName"`
	ServiceCode  string    `json:"serviceCode"`
	ServiceName  string    `json:"serviceName"`
	Path         string    `json:"path"`
	OriginalPath string    `json:"originalPath"`
	Method       string    `json:"method"`
	JobID        string    `json:"jobID"`
	Module       string    `json:"module"`
	Type         string    `json:"type"`
	ToPath       string    `json:"toPath"`
	ToPathName   string    `json:"toPathName"`
	Source       string    `json:"source"`
	StatusCode   int       `json:"statusCode"`
	ElapsedTime  int       `json:"elapsedTime"`
	Body         string    `json:"body"`
	Header       any       `json:"header"`
	Message      string    `json:"message"`
	RequestId    string    `json:"requestId"`
	Sql          string    `json:"sql"`
	Topic        string    `json:"topic"`
	Service      string    `json:"requested-service"`
	Protocol     string    `json:"protocol"`
	EntryModule  string    `json:"entryModule"`
	ErrorMessage any       `json:"errorMessage"`
	EventName    string    `json:"name"`
	PodName      string    `json:"pod_name"`
	HostName     string    `json:"host_name"`
	Host         string    `json:"host"`
	Limit        int       `json:"limit"`
}
