package config

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var (
	ErrConfigNotFound = errors.New("config file not found")
)

type Config struct {
	Version  string         `yaml:"version"`
	DataDir  string         `yaml:"data_directory"`
	Profiles []Profile      `yaml:"profiles"`
	AI       AIConfig       `yaml:"ai"`
	Tools    ToolsConfig    `yaml:"tools"`
	Sysadmin SysadminConfig `yaml:"sysadmin"`
}

type Profile struct {
	Name         string   `yaml:"name"`
	Type         string   `yaml:"type"`
	Hosts        []string `yaml:"hosts"`
	User         string   `yaml:"user"`
	Port         int      `yaml:"port"`
	KeyFile      string   `yaml:"key_file"`
	ForwardAgent bool     `yaml:"forward_agent"`
}

type AIConfig struct {
	Provider  string            `yaml:"provider"`
	Endpoint  string            `yaml:"endpoint"`
	Model     string            `yaml:"model"`
	MaxTokens int               `yaml:"max_tokens"`
	APIKeys   map[string]string `yaml:"api_keys"`
}

type ToolsConfig struct {
	AutoApprove   []string `yaml:"auto_approve"`
	ManualApprove []string `yaml:"manual_approve"`
}

type SysadminConfig struct {
	Resources ResourceConfig `yaml:"resources"`
	Logs      LogsConfig     `yaml:"logs"`
	Services  ServicesConfig `yaml:"services"`
}

type ResourceConfig struct {
	CPUWarning     float64 `yaml:"cpu_warning"`
	CPUCritical    float64 `yaml:"cpu_critical"`
	MemoryWarning  float64 `yaml:"memory_warning"`
	MemoryCritical float64 `yaml:"memory_critical"`
}

type LogsConfig struct {
	DefaultSources []string `yaml:"default_sources"`
}

type ServicesConfig struct {
	SupportedManagers []string `yaml:"supported_managers"`
}

func Load() (*Config, error) {
	cfgPath := filepath.Join(getConfigDir(), "config.yaml")

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrConfigNotFound
		}
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) validate() error {
	if c.Version == "" {
		return errors.New("version is required")
	}

	if c.AI.Provider == "" {
		c.AI.Provider = "local"
	}

	if c.AI.MaxTokens == 0 {
		c.AI.MaxTokens = 4096
	}

	if len(c.Sysadmin.Logs.DefaultSources) == 0 {
		c.Sysadmin.Logs.DefaultSources = []string{
			"/var/log/syslog",
			"/var/log/auth.log",
		}
	}

	if len(c.Sysadmin.Services.SupportedManagers) == 0 {
		c.Sysadmin.Services.SupportedManagers = []string{"systemd"}
	}

	return nil
}

func getConfigDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "~/.config/smoothssh"
	}
	return filepath.Join(home, ".config", "smoothssh")
}

func getDataDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "~/.local/share/smoothssh"
	}
	return filepath.Join(home, ".local", "share", "smoothssh")
}
