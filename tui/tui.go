package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mreeves/smoothssh/ai"
	"github.com/mreeves/smoothssh/config"
	"github.com/mreeves/smoothssh/model"
	"github.com/mreeves/smoothssh/ssh"
)

type Model struct {
	config   *config.Config
	quitting bool
	session  *ai.Session
	client   *ssh.Client
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
		case "ctrl+r":
			return m.initSession()
		}
	}
	return m, nil
}

func (m *Model) initSession() (tea.Model, tea.Cmd) {
	if len(m.config.Profiles) == 0 {
		return m, nil
	}

	profile := m.config.Profiles[0]

	cfg := &model.SSHConfig{
		Hostname:     profile.Hosts[0],
		User:         profile.User,
		Port:         profile.Port,
		KeyFile:      profile.KeyFile,
		ForwardAgent: profile.ForwardAgent,
	}

	client, err := ssh.New(cfg)
	if err != nil {
		return m, nil
	}

	if err := client.Connect(); err != nil {
		return m, nil
	}

	m.client = client

	sshConn := &ai.SSHConnection{
		Hostname: profile.Hosts[0],
		User:     profile.User,
		Client:   client,
	}

	m.session = ai.NewSession(sshConn)

	return m, nil
}

func (m *Model) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}

	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("63")).
		Render("SmoothSSH - SSH AI Assistant TUI")

	subheader := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Render(fmt.Sprintf("Profile: %s | AI: %s", m.config.Profiles[0].Name, m.config.AI.Provider))

	content := strings.Builder{}
	content.WriteString(header + "\n")
	content.WriteString(subheader + "\n")
	content.WriteString(strings.Repeat("=", 40) + "\n\n")

	if m.session == nil {
		content.WriteString("Press Ctrl+R to connect and initialize AI session\n")
	} else {
		content.WriteString("✓ Session connected - AI assistant ready\n")
	}

	content.WriteString("\n")
	content.WriteString("Key Bindings:\n")
	content.WriteString("  q, esc, Ctrl+C - Quit\n")
	content.WriteString("  Ctrl+R         - Reset/Reconnect\n")

	return content.String()
}

func Run(cfg *config.Config) error {
	p := tea.NewProgram(New(cfg), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("failed to run program: %w", err)
	}
	return nil
}
