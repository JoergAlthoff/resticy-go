package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Backup       BackupConfig    `yaml:"backup,omitempty"`
	Forget       ForgetConfig    `yaml:"forget,omitempty"`
	Parent       ParentConfig    `yaml:"parent,omitempty"`
	Check        CheckConfig     `yaml:"check,omitempty"`
	Stats        StatsConfig     `yaml:"stats,omitempty"`
	Snapshots    SnapshotsConfig `yaml:"snapshots,omitempty"`
	List         ListConfig      `yaml:"list,omitempty"`
	Prune        PruneConfig     `yaml:"prune,omitempty"`
	BackupLog    string          `yaml:"backup_log"`
	ForgetLog    string          `yaml:"forget_log,omitempty"`
	StatusLog    string          `yaml:"status_log,omitempty"`
	PruneLog     string          `yaml:"prune_log,omitempty"`
	InfoLog      string          `yaml:"info_log,omitempty"`
	Debug        bool            `yaml:"debug,omitempty"`
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

func (cfg *AppConfig) ApplyDefaults() {
	cfg.Parent.ApplyDefaults()
	cfg.Stats.ApplyDefaults()
	cfg.Snapshots.ApplyDefaults()
	cfg.List.ApplyDefaults()
	cfg.Backup.ApplyDefaults()
	cfg.Forget.ApplyDefaults()
	cfg.Check.ApplyDefaults()
	cfg.Prune.ApplyDefaults()

	if cfg.BackupLog == "" {
		cfg.BackupLog = "/var/log/restic_backup.log"
	}
	if cfg.ForgetLog == "" {
		cfg.ForgetLog = "/var/log/restic_forget.log"
	}
	if cfg.StatusLog == "" {
		cfg.StatusLog = "/var/log/restic_status.log"
	}
	if cfg.PruneLog == "" {
		cfg.PruneLog = "/var/log/restic_prune.log"
	}
	if cfg.InfoLog == "" {
		cfg.InfoLog = "/var/log/restic_info.log"
	}

	if err := cfg.Parent.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "Invalid parent config: %v\n", err)
		os.Exit(1)
	}
}

func (cfg *AppConfig) Validate() error {
	if err := cfg.Backup.Validate(); err != nil {
		return fmt.Errorf("invalid backup config: %w", err)
	}
	if err := cfg.Forget.Validate(); err != nil {
		return fmt.Errorf("invalid forget config: %w", err)
	}
	if err := cfg.Check.Validate(); err != nil {
		return fmt.Errorf("invalid check config: %w", err)
	}
	if err := cfg.Parent.Validate(); err != nil {
		return fmt.Errorf("invalid parent config: %w", err)
	}
	if err := cfg.Stats.Validate(); err != nil {
		return fmt.Errorf("invalid stats config: %w", err)
	}
	if err := cfg.Snapshots.Validate(); err != nil {
		return fmt.Errorf("invalid snapshots config: %w", err)
	}
	if err := cfg.List.Validate(); err != nil {
		return fmt.Errorf("invalid list config: %w", err)
	}
	if err := cfg.Prune.Validate(); err != nil {
		return fmt.Errorf("invalid prune config: %w", err)
	}
	return nil
}

// BuildFlags implements the ConfigSection interface.
// AppConfig itself does not contribute restic CLI flags directly,
// but delegates that responsibility to its sub-configurations.
func (cfg *AppConfig) BuildFlags() []string {
	return nil
}

var _ ConfigSection = (*AppConfig)(nil)
