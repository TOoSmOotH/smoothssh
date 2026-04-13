package ai

import (
	"fmt"
	"strings"
)

type Tool struct {
	Name        string
	Description string
	Exec        func(args []string) (string, error)
}

type Session struct {
	Messages   []Message
	Tools      []Tool
	Connection *SSHConnection
}

type SSHConnection struct {
	Hostname string
	User     string
	Client   interface{}
}

func NewSession(conn *SSHConnection) *Session {
	return &Session{
		Messages:   []Message{},
		Tools:      defaultTools(),
		Connection: conn,
	}
}

func defaultTools() []Tool {
	return []Tool{
		{
			Name:        "ls",
			Description: "List directory contents",
			Exec: func(args []string) (string, error) {
				return runSSHCommand("ls", args)
			},
		},
		{
			Name:        "pwd",
			Description: "Print working directory",
			Exec: func(args []string) (string, error) {
				return runSSHCommand("pwd", args)
			},
		},
		{
			Name:        "cat",
			Description: "Display file contents",
			Exec: func(args []string) (string, error) {
				if len(args) == 0 {
					return "", fmt.Errorf("cat requires a file argument")
				}
				return runSSHCommand("cat", args)
			},
		},
		{
			Name:        "grep",
			Description: "Search text patterns",
			Exec: func(args []string) (string, error) {
				if len(args) < 2 {
					return "", fmt.Errorf("grep requires pattern and file arguments")
				}
				return runSSHCommand("grep", args)
			},
		},
		{
			Name:        "ps",
			Description: "List processes",
			Exec: func(args []string) (string, error) {
				return runSSHCommand("ps", args)
			},
		},
		{
			Name:        "df",
			Description: "Disk space usage",
			Exec: func(args []string) (string, error) {
				return runSSHCommand("df", args)
			},
		},
		{
			Name:        "top",
			Description: "Display system resources",
			Exec: func(args []string) (string, error) {
				return runSSHCommand("top", args)
			},
		},
		{
			Name:        "journalctl",
			Description: "View systemd logs",
			Exec: func(args []string) (string, error) {
				return runSSHCommand("journalctl", args)
			},
		},
	}
}

func runSSHCommand(cmd string, args []string) (string, error) {
	if args == nil {
		args = []string{}
	}
	fullCmd := cmd + " " + strings.Join(args, " ")
	return "cmd: " + fullCmd, nil
}

func (s *Session) AddMessage(role, content string) {
	s.Messages = append(s.Messages, Message{Role: role, Content: content})
}

func (s *Session) Chat(aiClient *Client, content string) (string, error) {
	s.AddMessage("user", content)

	resp, err := aiClient.Chat(s.Messages)
	if err != nil {
		return "", err
	}

	s.AddMessage("assistant", resp)
	return resp, nil
}
