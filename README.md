# SmoothSSH - SSH AI Assistant TUI

A rich TUI SSH client with integrated AI assistant capabilities, focused on system administration tasks on remote Linux/Unix servers.

## Features

- **SSH Connection Management**: YAML-based configuration with connection pooling
- **AI Assistant**: Local models (Ollama) + cloud providers (OpenAI, Anthropic, Groq)
- **System Admin Tools**: Resource monitor, log viewer, process manager, service manager
- **Workflow Automation**: YAML workflows and Agent Skills support

## Installation

```bash
go install github.com/mreeves/smoothssh/cmd/smoothssh@latest
```

## Configuration

Create `~/.config/smoothssh/config.yaml`:

```yaml
version: "1.0"
data_directory: ~/.local/share/ssh-ai

profiles:
  - name: production-servers
    type: ssh
    hosts:
      - web01.prod.example.com
    user: admin
    port: 22
    key_file: ~/.ssh/id_ed25519
    forward_agent: true

ai:
  provider: local
  endpoint: http://localhost:11434
  model: qwen3:30b
  max_tokens: 4096

tools:
  permissions:
    auto_approve: [view, ls, grep, view_file]
    manual_approve: [bash, write, edit, run_service]
```

## Usage

```bash
smoothssh
```

## KeyBindings

- `q`, `esc`, `ctrl+c` - Quit
- `h`, `left` - Previous view
- `l`, `right` - Next view
- `ctrl+p` - Profile selector
- `ctrl+a` - AI assistant

## Status

**Phase 1 (Foundation)**: In progress
- ✅ Project structure
- ✅ Config system
- ✅ SSH client
- ✅ Basic TUI
- ⏳ AI integration
- ⏳ File browser
- ⏳ Process viewer

## License

MIT
