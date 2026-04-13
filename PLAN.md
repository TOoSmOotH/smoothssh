# SSH AI Assistant TUI - Plan

## Project Overview

A rich TUI SSH client with integrated AI assistant capabilities, inspired by Crush/OpenCode but focused on system administration tasks on remote Linux/Unix servers.

## Current Status

- **Mode:** Build phase
- **Target:** Fresh implementation (not based on existing codebase)

## Technical Stack

- **Language:** Go
- **TUI Framework:** Bubble Tea v2
- **SSH Library:** golang.org/x/crypto/ssh
- **AI Integration:** Crush-compatible agent architecture
- **Storage:** SQLite for sessions and history
- **Editor:** Bubble Tea-based terminal editor

## Requirements (from user)

### Primary Goals
1. System administration assistant
2. Automated workflows support

### Interaction Model
- Hybrid approach (both conversational TUI and inline command assistance)

### Platform Support
- Linux/Unix servers (primary focus)

### Configuration
- YAML config only (not SSH config compatibility）

### AI Preferences
- **Primary:** Local models (Ollama/LM Studio)
- **Secondary:** OpenAI, Anthropic, Groq support

### Key Sysadmin Tools Needed
1. System resource monitor (CPU, memory, disk, network)
2. Log viewer with tail/filter
3. Remote process viewer/manager
4. File browser/manager
5. Service management UI
6. Network diagnostics

### Automation
- Both YAML workflows and Agent Skills format support

## Core Components

### 1. Connection Management
- YAML-based server configuration
- Connection pooling and multiplexing
- SSH agent forwarding
- Zeroconf/DNS-SD discovery
- Connection health monitoring

### 2. Hybrid AI Interface
- **Chat Mode:** Full TUI with Bubble Tea interface
- **Command Mode:** Inline AI assistance for specific tasks
- **Context-Aware Tools:** AI can interact with remote filesystem, execute commands
- **Permission System:** User approval before tool execution

### 3. AI Agent Capabilities
- Multi-provider support (OpenAI, Anthropic, Groq, OpenRouter, Local models)
- Session persistence and management
- Tool permissions system
- Custom skills support (Agent Skills standard + YAML workflows)
- Context window management with auto-compaction
- Model switching per session

### 4. System Administration Tools

#### Resource Monitor
- Live CPU/memory/disk/network usage
- Historical graphs
- Alerts and thresholds
- Per-process resource tracking

#### Log Viewer
- Real-time tail with filtering
- Log rotation detection
- Pattern matching and grep
- Multiple log source support
- Export capabilities

#### File Browser/Manager
- Directory navigation with fuzzy search
- Integrated editor (Bubble Tea-based)
- File permissions management
- Upload/download
- SFTP integration

#### Process Manager
- Live process listing with filtering
- Process tree visualization
- Kill/signal sending
- Process details view
- Resource usage per process

#### Service Manager
- Service listing (systemd/init/elk)
- Start/stop/restart controls
- Service status monitoring
- Journal viewing
- Enable/disable capabilities

#### Network Diagnostics
- Connection status viewed
- Port scanning
- Network interface monitoring
- Traceroute/latency tools
- DNS resolution tools

### 5. Workflow Automation

#### YAML Workflows
- Task definitions with variables
- Multi-server execution
- Conditional logic
- Error handling
- Result aggregation

#### Agent Skills
- Open standard compatibility
- Local/skill discovery
- Shared skill libraries
- Version control integration

## Configuration Example

```yaml
version: "1.0"
data_directory: ~/.local/share/ssh-ai

profiles:
  - name: production-servers
    type: ssh
    hosts:
      - web01.prod.example.com
      - web02.prod.example.com
      - db01.prod.example.com
    user: admin
    port: 22
    key_file: ~/.ssh/id_ed25519
    forward_agent: true
    
  - name: development-local
    type: ssh
    hosts:
      - localhost:2222
    user: devuser
    port: 2222
    
ai:
  provider: local
  endpoint: http://localhost:11434
  model: qwen3:30b
  max_tokens: 4096
  
  api_keys:
    openai: $OPENAI_API_KEY
    anthropic: $ANTHROPIC_API_KEY
    groq: $GROQ_API_KEY
    
tools:
  permissions:
    auto_approve:
      - view
      - ls
      - grep
      - view_file
      
    manual_approve:
      - bash
      - write
      - edit
      - run_service

sysadmin:
  resources:
    cpu_warning: 80
    cpu_critical: 95
    memory_warning: 85
    memory_critical: 95
    
  logs:
    default_sources:
      - /var/log/syslog
      - /var/log/auth.log
      - /var/log/daemon.log
      
  services:
    supported_managers:
      - systemd
      - openbsdrc
      - s6
```

## Implementation Phases

### Phase 1: Foundation (Core SSH + AI)
1. Project structure and config system
2. SSH connection management
3. Basic Bubble Tea TUI layout
4. Local model integration (Ollama)
5. Session management
6. File browsing
7. Process viewer

### Phase 2: Intelligence (AI Enhancement)
1. Multi-provider support
2. Tool permissions system
3. Agent Skills support
4. Log viewer with filtering
5. Resource monitor
6. Service management UI
7. Network diagnostics

### Phase 3: Automation
1. YAML workflow support
2. Multi-server execution
3. Custom command system
4. Results aggregation
5. Workflow history
6. Scheduling capabilities

### Phase 4: Polish & Advanced Features
1. Performance optimization
2. Advanced filtering
3. Custom themes
4. Keyboard shortcuts
5. Documentation
6. Plugin system

## Project Name Suggestions

- **Stream** (SSH + Terminal + Remote + AI + Model)
- **Remotely** (Remote + AI assistant)
- **Overthere** (SSH jargon for remote systems)
- **Remix** (Remote + AI + Mixed interface)

## Dependencies to Consider

### Core Libraries
- `charm.land/bubbletea/v2` - TUI framework
- `golang.org/x/crypto/ssh` - SSH client
- `sqlite.org` - Database
- `github.com/charmbracelet/lipgloss` - Styling
- `github.com/charmbracelet/bubbles` - UI components

### AI Integration
- Crush/OpenCode agent architecture patterns
- OpenAI/Anthropic SDKs
- Local model HTTP clients

### Sysadmin Features
- `github.com/shirou/gopsutil` - System stats
- Terminal rendering libraries for log viewing

## Open Questions

1. Should this be a standalone tool or integrated with existing SSH tooling?
2. How should we handle multi-server operations (parallel vs. sequential)?
3. Do we need persistent sessions or connection sharing?
4. Should we include a built-in terminal emulator for full compatibility?
5. What authentication methods beyond SSH keys should be supported?
