# SmoothSSH - Implementation Notes

## Project Status

### Phase 1: Foundation (Core SSH + AI) - In Progress

#### Completed
- ✅ Project structure and module setup
- ✅ Configuration system (YAML-based)
- ✅ SSH client with key/password auth
- ✅ Basic Bubble Tea TUI layout
- ✅ Go module and build system

#### In Progress
- 🔄 TUI components (file browser, process viewer)

#### Next Steps
- AI integration (Ollama)
- Session management
- File browser UI
- Process viewer UI

## Development Setup

### Prerequisites
- Go 1.26+
- Ollama (for local AI models)

### Quick Start

```bash
# Setup config
./scripts/setup.sh

# Build
./scripts/build.sh

# Run
./scripts/quickstart.sh
```

## Code Structure

```
smoothssh/
├── cmd/
│   └── smoothssh/
│       └── main.go          # Entry point
├── config/                   # Configuration handling
│   ├── config.go
│   └── doc.go
├── model/                    # Data models
│   └── ssh.go
├── ssh/                      # SSH connection management
│   ├── client.go
│   └── doc.go
├── tui/                      # Terminal UI
│   ├── tui.go
│   └── components/           # TUI components
├── session/                  # AI session management
├── ai/                       # AI provider integration
├── db/                       # SQLite storage
├── workflow/                 # Automation workflows
├── scripts/                  # Build/Setup scripts
└── PLAN.md                   # Project plan
```

## Configuration

Config location: `~/.config/smoothssh/config.yaml`

Example config provided in root directory.

## Next Milestones

1. Complete TUI component library
2. Ollama integration for local AI
3. File browser with SFTP support
4. Process viewer with filtering
5. Log viewer with tail/filter
6. Session persistence

## Testing

```bash
go test ./...
```

## License

MIT © 2026 Mike Reeves
