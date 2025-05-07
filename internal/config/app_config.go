package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Backup       BackupConfig `yaml:"backup,omitempty"`
	Forget       ForgetConfig `yaml:"forget,omitempty"`
	Parent       ParentConfig `yaml:"parent,omitempty"`
	Check        CheckConfig  `yaml:"check,omitempty"`
	BackupLog    string       `yaml:"backup_log"`
	ForgetLog    string       `yaml:"forget_log,omitempty"`
	SnapshotsLog string       `yaml:"snapshots_log,omitempty"`
	StatusLog    string       `yaml:"status_log,omitempty"`
	Debug        bool         `yaml:"debug,omitempty"`
}

func Load(path string) (*AppConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %w", err)
	}
	var cfg AppConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("invalid YAML format: %w", err)
	}
	return &cfg, nil
}

func ApplyDefaults(cfg *AppConfig) {
	cfg.Parent.ApplyDefaults()

	if cfg.BackupLog == "" {
		cfg.BackupLog = "/var/log/restic_backup.log"
	}
	if cfg.ForgetLog == "" {
		cfg.ForgetLog = "/var/log/restic_forget.log"
	}
	if cfg.SnapshotsLog == "" {
		cfg.SnapshotsLog = "/var/log/restic_snapshots.log"
	}
	if cfg.StatusLog == "" {
		cfg.StatusLog = "/var/log/restic_status.log"
	}

	if err := cfg.Parent.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "Invalid parent config: %v\n", err)
		os.Exit(1)
	}
}
