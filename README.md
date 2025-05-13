# resticy-go

**resticy-go** is a modular command-line tool written in Go that simplifies working with [restic](https://restic.net/) backups. It provides a structured YAML-based configuration system and convenient commands for managing backup repositories.

## Features

- Run restic commands with predefined configuration
- Manage snapshots, check repository status, and apply policies
- Modular subcommands for future extensibility
- Clean CLI interface with error handling and reporting
- Local mirroring to GitHub and private Gitea instance

## Requirements

- Go 1.21 or higher
- A working [restic](https://restic.net/) installation
- A `config.yaml` file tailored to your environment

## Getting Started

1. Build the binary:

```bash
go build -o resticy ./cmd/resticy
```

2. Run a command:

```bash
./resticy snapshots --config=config.yaml
```

## License

This project is licensed under the [GNU Affero General Public License v3.0](LICENSE).

## Status

This project is under active development. Contributions are welcome.