package config

import "fmt"

type BackupConfig struct {
	Exclude []string `yaml:"exclude,omitempty"`
	// ggf. weitere Felder folgen
}

func (c *AppConfig) BuildExcludeFlags() []string {
	var flags []string
	for _, pattern := range c.Backup.Exclude {
		flags = append(flags, fmt.Sprintf("--exclude=%s", pattern))
	}
	return flags
}
