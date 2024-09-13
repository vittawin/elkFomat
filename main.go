package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

//var data []string
//var rows = make(map[string][]LogStruct) //keys is jobId of log
//var keys = make([]string, 0)
//var dataLogList []list.Item

func main() {
	setInput()
	p := tea.NewProgram(&MyModel{})
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
