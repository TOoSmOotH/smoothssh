# Contributing to SmoothSSH

Thank you for your interest in contributing to SmoothSSH! This document provides guidelines for contributing.

## Getting Started

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/your-feature`
3. Make your changes
4. Run tests: `go test ./...`
5. Submit a pull request

## Development Setup

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/smoothssh.git
cd smoothssh

# Install dependencies
go mod download

# Build
go build ./cmd/smoothssh
```

## Code Style

- Follow Go conventions (gofmt)
- Use meaningful variable names
- Add comments for complex logic
- Keep functions focused and small

## Commit Messages

Use conventional commits:
- `feat: add new feature`
- `fix: fix bug`
- `docs: update documentation`
- `refactor: refactor code`
- `test: add tests`

## Pull Request Process

1. Update README.md if needed
2. Update documentation
3. Add tests for new functionality
4. Ensure CI passes
5. Address review feedback

## Questions?

Open an issue or contact the maintainers.
