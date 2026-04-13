package components

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ProcessViewer struct {
	processes []ProcessEntry
	selected  int
}

type ProcessEntry struct {
	PID     int
	USER    string
	CPU     float64
	MEM     float64
	COMMAND string
}

func NewProcessViewer() *ProcessViewer {
	return &ProcessViewer{
		processes: []ProcessEntry{
			{PID: 1, USER: "root", CPU: 0.1, MEM: 1.2, COMMAND: "/sbin/init"},
			{PID: 1234, USER: "user", CPU: 5.5, MEM: 10.3, COMMAND: "bash"},
			{PID: 5678, USER: "user", CPU: 2.1, MEM: 5.7, COMMAND: "python3 app.py"},
		},
	}
}

func (m *ProcessViewer) Init() tea.Cmd {
	return nil
}

func (m *ProcessViewer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "m":
			return m, tea.Batch(func() tea.Msg {
				return ReturnToMainMsg{}
			})
		case "q", "esc":
			return m, tea.Quit
		case "j", "down":
			if m.selected < len(m.processes)-1 {
				m.selected++
			}
		case "k", "up":
			if m.selected > 0 {
				m.selected--
			}
		}
	}
	return m, nil
}

func (m *ProcessViewer) View() string {
	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("63")).
		Render("Process Viewer") + "\n\n"

	list := ""
	for i, p := range m.processes {
		selected := "  "
		if i == m.selected {
			selected = "> "
		}
		row := fmt.Sprintf("%-8d %-12s %-8.1f %-8.1f %s",
			p.PID, p.USER, p.CPU, p.MEM, p.COMMAND)
		list += lipgloss.NewStyle().
			Width(80).
			Foreground(lipgloss.Color("240")).
			Render(selected + row) + "\n"
	}

	return header + list + "\n" +
		"Key Bindings:\n" +
		"  m - Return to main view\n" +
		"  q - Quit\n" +
		"  j/k, up/down - Navigate\n"
}
