package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MyModel struct {
	list     list.Model
	secModel SecModel
	err      error
}

func (m *MyModel) Init() tea.Cmd {
	return nil
}

func (m *MyModel) initList(w int) {
	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), w, 50)
	m.list.Title = "Elk logs List"
	m.list.SetItems(dataLogList)

	a := SecModel{}
	a.list = list.New([]list.Item{}, list.NewDefaultDelegate(), w, 50)
	a.list.Title = "Elkkkkkkkkkkku..."
}

func (m *MyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.initList(msg.Width)
	case tea.KeyMsg:
		switch msg.String() {
		case "g":
			m.list.SetItems(dataLogList)
		case "enter":
			m.SetLog()
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *MyModel) View() string {
	if m.secModel.jobId != "" {
		return m.secModel.View()
	}
	return lipgloss.Place(
		100,
		50,
		lipgloss.Left,
		lipgloss.Center,
		m.list.View(),
	)
}

func (m *MyModel) SetLog() {
	selectedItem := m.list.SelectedItem()
	selectedLog := selectedItem.(Log)
	//newDesc, err := parseLogBody(selectedLog.data)
	//if err != nil {
	//	panic(err)
	//}

	selectedLog.description = "aaaaaaaaaaa\n" + "wwwwwww"
	m.secModel.list.SetItems([]list.Item{selectedLog})
}
