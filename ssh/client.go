package ssh

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"

	"github.com/mreeves/smoothssh/model"
)

type Client struct {
	config *model.SSHConfig
	client *ssh.Client
}

func New(cfg *model.SSHConfig) (*Client, error) {
	c := &Client{config: cfg}
	return c, nil
}

func (c *Client) Connect() error {
	addr := fmt.Sprintf("%s:%d", c.config.Hostname, c.config.Port)

	config := &ssh.ClientConfig{
		User:            c.config.User,
		Timeout:         30,
		HostKeyCallback: nil,
	}

	if c.config.KeyFile != "" {
		keyPath, _ := filepath.Abs(c.config.KeyFile)
		keyBytes, err := os.ReadFile(keyPath)
		if err != nil {
			return fmt.Errorf("failed to read key file: %w", err)
		}

		signer, err := ssh.ParsePrivateKey(keyBytes)
		if err != nil {
			return fmt.Errorf("failed to parse private key: %w", err)
		}

		config.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	} else {
		config.Auth = []ssh.AuthMethod{ssh.PasswordCallback(c.getPassword)}
	}

	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}

	c.client = client
	return nil
}

func (c *Client) getPassword() (string, error) {
	fmt.Print("Enter password: ")
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	fmt.Println()
	return string(bytePassword), nil
}

func (c *Client) Close() error {
	if c.client != nil {
		return c.client.Close()
	}
	return nil
}

func (c *Client) Execute(cmd string) (string, error) {
	session, err := c.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", err
	}

	return string(output), nil
}
