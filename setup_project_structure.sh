#!/bin/bash

# Create project folder structure for resticy-go
mkdir -p cmd/resticy
mkdir -p internal/config
mkdir -p internal/logging
mkdir -p internal/restic

# Create placeholder files
touch cmd/resticy/main.go
touch internal/config/config.go
touch internal/config/backup.go
touch internal/config/forget.go
touch internal/config/check.go
touch internal/config/global.go
touch internal/logging/init.go
touch internal/restic/restic.go


# Make script executable
chmod +x setup_project_structure.sh

echo "Project structure initialized."