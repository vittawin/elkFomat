package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int

var (
	columnStyle  = lipgloss.NewStyle().Padding(0, 2, 0, 2).Foreground(lipgloss.Color("241"))
	focusedStyle = lipgloss.NewStyle().Padding(0, 2, 0, 0).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
)

type MyModel struct {
	focused  status
	lists    []list.Model
	logData  string
	err      error
	loaded   bool
	viewport viewport.Model // Viewport for long log data
	list.Model
}

func (m *MyModel) Init() tea.Cmd {
	return nil
}

func (m *MyModel) initList(h, w int) {
	defaultDelegate := list.NewDefaultDelegate()
	logListDelegate := defaultDelegate
	logListDelegate.SetHeight(6)
	defaultList := list.New([]list.Item{}, logListDelegate, w/2, h*3/4)

	m.viewport = viewport.New(w/2, h)                        // Viewport size
	m.viewport.SetContent("Select a log to view details...") // Default content

	m.lists = []list.Model{defaultList}
	m.lists[0].Title = "Elk logs List"
	m.lists[0].SetItems(dataLogList)
}

func (m *MyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.initList(msg.Height, msg.Width)
			m.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "g":
			m.lists[0].SetItems(dataLogList)
		case "j", "down":
			m.SetLog(true)
		case "k", "up":
			m.SetLog(false)
		case "ctrl+d":
			m.viewport.LineDown(5) // Scroll down
		case "ctrl+u":
			m.viewport.LineUp(5) // Scroll up
		}
	}

	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m *MyModel) View() string {
	if m.loaded {
		logListView := m.lists[0].View()
		logDataView := m.viewport.View() // Render the viewport

		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			focusedStyle.Render(logListView),
			columnStyle.Render(logDataView),
		)
	} else {
		return "loading..."
	}
}

func (m *MyModel) setDefaultLogData() {
	if len(dataLogList) == 0 {
		return
	}

	firstDataLog := dataLogList[0].(Log)
	logsData, err := parseLogBody(firstDataLog.data)
	if err != nil {
		panic(err)
	}

	m.lists[1].SetItems([]list.Item{
		Log{
			jobId:       firstDataLog.jobId,
			title:       firstDataLog.jobId,
			description: logsData,
		},
	})
}

func (m *MyModel) SetLog(isDown bool) {
	selectedItem := m.SelectedItem(isDown)
	selectedLog := selectedItem.(Log)

	logsData, err := parseLogBody(selectedLog.data)
	if err != nil {
		panic(err)
	}

	m.logData = logsData
	m.viewport.SetContent(logsData) // Update viewport content
}

func (m *MyModel) SelectedItem(isDown bool) list.Item {
	i := m.lists[0].Index()

	items := m.lists[0].VisibleItems()
	itemsAmount := len(items)
	if i < 0 || len(items) == 0 || len(items) <= i {
		return nil
	}

	if isDown {
		i++
		if i >= itemsAmount {
			return items[i-1]
		}
	} else {
		i--
		if i < 0 {
			return items[0]
		}
	}
	return items[i]
}
