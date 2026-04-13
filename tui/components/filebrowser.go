package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/lipgloss"
)

type FileBrowser struct {
	picker filepicker.Model
}

func NewFileBrowser() *FileBrowser {
	fp := filepicker.New()
	fp.ShowHidden = false

	return &FileBrowser{
		picker: fp,
	}
}

func (m *FileBrowser) Init() tea.Cmd {
	return nil
}

func (m *FileBrowser) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "m":
			return m, tea.Batch(func() tea.Msg {
				return ReturnToMainMsg{}
			})
		case "q", "esc":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.picker, cmd = m.picker.Update(msg)
	return m, cmd
}

func (m *FileBrowser) View() string {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("63")).
		Render("File Browser") + "\n\n" +
		m.picker.View() + "\n\n" +
		"Key Bindings:\n" +
		"  m - Return to main view\n" +
		"  q - Quit\n"
}

type ReturnToMainMsg struct{}
