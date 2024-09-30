package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int

var (
	columnStyle  = lipgloss.NewStyle().Padding(0, 2, 0, 2)
	focusedStyle = lipgloss.NewStyle().Padding(0, 2, 0, 0).
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

	logDataDelegate := defaultDelegate
	logDataDelegate.SetHeight(100)
	logDataList := list.New([]list.Item{}, logDataDelegate, w/2, h)
	logDataList.SetShowHelp(false)

	m.lists = []list.Model{defaultList, logDataList}
	m.lists[0].Title = "Elk logs List"
	m.lists[0].SetItems(dataLogList)

	m.lists[1].Title = "log Data"
	m.setDefaultLogData()
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
			m.lists[m.focused].SetItems(dataLogList)
			m.lists[1].SetItems([]list.Item{})
		case "j", "down":
			m.SetLog(true)
		case "k", "up":
			m.SetLog(false)
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

func (m *MyModel) setDefaultLogData() {
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

	//m.lists[0].SetItems(filterLog(selectedLog.jobId))
	m.logData = logsData

	m.lists[1].SetItems([]list.Item{Log{
		jobId:       selectedLog.jobId,
		title:       selectedLog.jobId,
		description: logsData,
	}})
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
func filterLog(jobId string) []list.Item {
	var filteredLogs []list.Item
	for _, log := range dataLogList {
		if log.(Log).jobId == jobId {
			filteredLogs = append(filteredLogs, log)
		}
	}
	return filteredLogs
}
