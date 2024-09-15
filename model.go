package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int

var (
	columnStyle  = lipgloss.NewStyle().Padding(1, 2, 0, 5)
	focusedStyle = lipgloss.NewStyle().Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
)

type MyModel struct {
	focused  status
	lists    []list.Model
	logData  string
	secModel SecModel
	err      error
	loaded   bool //use this for wait until it finish all setting items -> remove this will error
}

func (m *MyModel) Init() tea.Cmd {
	return nil
}

func (m *MyModel) initList(w int) {
	defaultDelegate := list.NewDefaultDelegate()
	defaultList := list.New([]list.Item{}, defaultDelegate, w, 50)

	logDataDelegate := defaultDelegate
	logDataDelegate.SetHeight(50)
	logDataList := list.New([]list.Item{}, logDataDelegate, w, 50)
	logDataList.SetShowHelp(false)

	m.lists = []list.Model{defaultList, logDataList}
	m.lists[0].Title = "Elk logs List"
	m.lists[0].SetItems(dataLogList)

	m.lists[1].Title = "log Data"
}

func (m *MyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.initList(msg.Width)
			m.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "g":
			m.lists[m.focused].SetItems(dataLogList)
			m.lists[1].SetItems([]list.Item{})
		case "enter":
			m.SetLog()
		}
	}
	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m *MyModel) View() string {
	if m.loaded {
		logListView := m.lists[0].View()
		logDataView := m.lists[1].View()
		switch m.focused {
		default:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				focusedStyle.Render(logListView),
				columnStyle.Render(logDataView),
			)
		}

	} else {
		return "loading..."
	}
}

func (m *MyModel) SetLog() {
	selectedItem := m.lists[0].SelectedItem()
	selectedLog := selectedItem.(Log)
	logsData, err := parseLogBody(selectedLog.data)
	if err != nil {
		panic(err)
	}

	m.lists[0].SetItems(filterLog(selectedLog.jobId))
	m.logData = logsData

	m.lists[1].SetItems([]list.Item{Log{
		jobId:       selectedLog.jobId,
		title:       selectedLog.jobId,
		description: logsData,
	}})
}

func filterLog(jobId string) []list.Item {
	var filteredLogs []list.Item
	for _, log := range dataLogList {
		if log.(Log).jobId == jobId {
			filteredLogs = append(filteredLogs, log)
		}
	}
	return filteredLogs
}
