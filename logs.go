package main

type Log struct {
	jobId       string
	data        string
	title       string
	description string
}

func (t Log) FilterValue() string {
	return t.jobId
}

func (t Log) Title() string {
	return t.jobId
}

func (t Log) Description() string {
	return t.description
}
