package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	setInput()
	p := tea.NewProgram(&MyModel{})
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
