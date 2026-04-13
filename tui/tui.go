package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mreeves/smoothssh/config"
)

type Model struct {
	config   *config.Config
	quitting bool
}

func New(config *config.Config) *Model {
	return &Model{config: config}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *Model) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}

	var sb strings.Builder
	sb.WriteString("SmoothSSH - SSH AI Assistant TUI\n")
	sb.WriteString("================================\n\n")
	sb.WriteString(fmt.Sprintf("Config: %s\n", m.config.Version))
	sb.WriteString(fmt.Sprintf("Profiles: %d\n", len(m.config.Profiles)))
	sb.WriteString(fmt.Sprintf("AI Provider: %s\n", m.config.AI.Provider))
	sb.WriteString("\n")
	sb.WriteString("Press q to quit\n")

	return sb.String()
}

func Run(cfg *config.Config) error {
	p := tea.NewProgram(New(cfg), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("failed to run program: %w", err)
	}
	return nil
}
