#!/bin/bash

# Create directories relative to the current directory
mkdir -p cmd/etl-ui
mkdir -p internal/ui/pages
mkdir -p internal/ui/components
mkdir -p pkg/config

# Create files
touch cmd/etl-ui/main.go
touch internal/ui/terminal.go
touch internal/ui/pages/main_menu.go
touch internal/ui/pages/create_job.go
touch internal/ui/pages/manage_jobs.go
touch internal/ui/pages/run_job.go
touch internal/ui/pages/settings.go
touch internal/ui/components/form.go
touch internal/ui/components/list.go
touch internal/ui/components/progress.go
touch internal/ui/components/logs.go
touch pkg/config/config.go

echo "Directory structure and files created successfully."
