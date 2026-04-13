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
	"github.com/mreeves/smoothssh/tui/components"
)

type ViewMode string

const (
	ViewMain       ViewMode = "main"
	ViewFile       ViewMode = "filebrowser"
	ViewProcess    ViewMode = "processviewer"
	ViewLog        ViewMode = "logviewer"
	ViewService    ViewMode = "servicemanager"
)

type Model struct {
	config      *config.Config
	quitting    bool
	session     *ai.Session
	client      *ssh.Client
	currentView ViewMode
	fileBrowser *components.FileBrowser
}

func New(config *config.Config) *Model {
	return &Model{
		config:      config,
		currentView: ViewMain,
	}
}

func (m *Model) Init() tea.Cmd {
	if m.fileBrowser == nil {
		m.fileBrowser = components.NewFileBrowser()
	}
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
		case "f":
			m.currentView = ViewFile
			return m, nil
		case "p":
			m.currentView = ViewProcess
			return m, nil
		case "l":
			m.currentView = ViewLog
			return m, nil
		case "s":
			m.currentView = ViewService
			return m, nil
		case "m":
			m.currentView = ViewMain
			return m, nil
		}
	}

	if m.fileBrowser != nil && m.currentView == ViewFile {
		model, cmd := m.fileBrowser.Update(msg)
		m.fileBrowser = model.(*components.FileBrowser)
		return m, cmd
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

	mainContent := strings.Builder{}
	mainContent.WriteString(header + "\n")
	mainContent.WriteString(subheader + "\n")
	mainContent.WriteString(strings.Repeat("=", 40) + "\n\n")

	if m.session == nil {
		mainContent.WriteString("Press Ctrl+R to connect and initialize AI session\n")
	} else {
		mainContent.WriteString("Session connected - AI assistant ready\n")
	}

	mainContent.WriteString("\n")
	mainContent.WriteString("Key Bindings:\n")
	mainContent.WriteString("  q, esc, Ctrl+C - Quit\n")
	mainContent.WriteString("  Ctrl+R         - Reset/Reconnect\n")

	if m.session != nil {
		mainContent.WriteString("\nSystem Admin Views:\n")
		mainContent.WriteString("  f - File Browser\n")
		mainContent.WriteString("  p - Process Viewer\n")
		mainContent.WriteString("  l - Log Viewer\n")
		mainContent.WriteString("  s - Service Manager\n")
	}

	switch m.currentView {
	case ViewFile:
		return m.getFileBrowserView()
	case ViewProcess:
		return m.getProcessViewerView()
	case ViewLog:
		return m.getLogViewerView()
	case ViewService:
		return m.getServiceManagerView()
	default:
		return mainContent.String()
	}
}

func (m *Model) getFileBrowserView() string {
	if m.fileBrowser != nil {
		return m.fileBrowser.View()
	}
	return "File Browser not initialized"
}

func (m *Model) getProcessViewerView() string {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("63")).
		Render("Process Viewer") + "\n\n" +
		"Live process listing coming soon...\n\n" +
		"Press m to return to main view\n"
}

func (m *Model) getLogViewerView() string {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("63")).
		Render("Log Viewer") + "\n\n" +
		"Tail and filter logs coming soon...\n\n" +
		"Press m to return to main view\n"
}

func (m *Model) getServiceManagerView() string {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("63")).
		Render("Service Manager") + "\n\n" +
		"Systemd service management coming soon...\n\n" +
		"Press m to return to main view\n"
}

func Run(cfg *config.Config) error {
	p := tea.NewProgram(New(cfg), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("failed to run program: %w", err)
	}
	return nil
}
