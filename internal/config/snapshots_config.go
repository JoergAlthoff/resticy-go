

package config

import "fmt"

// SnapshotsConfig holds configuration for the 'restic snapshots' command.
type SnapshotsConfig struct {
	Host     []string `yaml:"host,omitempty"`
	Path     []string `yaml:"path,omitempty"`
	Tag      []string `yaml:"tag,omitempty"`
	Compact  bool     `yaml:"compact,omitempty"`
	GroupBy  string   `yaml:"group-by,omitempty"`
	Latest   int      `yaml:"latest,omitempty"`
}

func (c *SnapshotsConfig) ApplyDefaults() {
	// No defaults currently required
}

func (c *SnapshotsConfig) Validate() error {
	if c.Latest < 0 {
		return fmt.Errorf("latest must not be negative")
	}
	return nil
}

func (c *SnapshotsConfig) BuildFlags() []string {
	var flags []string
	for _, h := range c.Host {
		flags = append(flags, "--host="+h)
	}
	for _, p := range c.Path {
		flags = append(flags, "--path="+p)
	}
	for _, t := range c.Tag {
		flags = append(flags, "--tag="+t)
	}
	if c.Compact {
		flags = append(flags, "--compact")
	}
	if c.GroupBy != "" {
		flags = append(flags, "--group-by="+c.GroupBy)
	}
	if c.Latest > 0 {
		flags = append(flags, fmt.Sprintf("--latest=%d", c.Latest))
	}
	return flags
}

var _ ConfigSection = (*SnapshotsConfig)(nil)