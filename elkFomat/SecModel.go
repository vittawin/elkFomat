package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SecModel struct {
	list        list.Model
	jobId       string
	title       string
	description string
}

func (m *SecModel) Init() tea.Cmd {
	return nil
}

func (m *SecModel) initList(w int) {
	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), w, 50)
	m.list.Title = "Elk logs List"
	m.list.SetItems(dataLogList)
}

func (m *SecModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.initList(msg.Width)
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *SecModel) View() string {
	return lipgloss.Place(
		100,
		50,
		lipgloss.Left,
		lipgloss.Center,
		m.list.View(),
	)
}
