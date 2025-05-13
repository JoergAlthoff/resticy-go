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

## Example Configuration

A sample `config.yaml` might look like this:

```yaml
repository: /mnt/restic/restic-beelink1
password_file: /etc/restic/.restic_password  # Path to your restic password file

log_file: /var/log/restic_full.log
log_level: info

backup:
  files-from: ""
  stdin: ""
  exclude:
    - /home/joerg/.cache
    - /home/joerg/.local/share/Trash
    - /var/tmp
  dry-run: false

check:
  read-data: false
  check-unused: true
```

Adjust paths and options to suit your environment. Do not include your password file in version control.

## Example Commands

Here are some example usages of `resticy`:

### Show snapshots

```bash
./resticy snapshots --config=config.yaml
```

### Check repository integrity

```bash
./resticy check --config=config.yaml
```

### Backup now

```bash
./resticy backup ~/data --config=config.yaml
```

### Show repository status

```bash
./resticy status --config=config.yaml
```

## License

This project is licensed under the [GNU Affero General Public License v3.0](LICENSE).

## Status

This project is under active development. Contributions are welcome.
