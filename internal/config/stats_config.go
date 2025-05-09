package config

import "fmt"

// StatsConfig holds configuration for the 'restic stats' command.
type StatsConfig struct {
	Mode        string   `yaml:"mode,omitempty"`
	Host        []string `yaml:"host,omitempty"`
	Path        []string `yaml:"path,omitempty"`
	Tag         []string `yaml:"tag,omitempty"`
	SnapshotIDs []string `yaml:"snapshot_ids,omitempty"`
}

func (c *StatsConfig) ApplyDefaults() {
	if c.Mode == "" {
		c.Mode = "restore-size"
	}
}

func (c *StatsConfig) Validate() error {
	validModes := map[string]bool{
		"restore-size":      true,
		"files-by-contents": true,
		"raw-data":          true,
		"blobs-per-file":    true,
	}
	if !validModes[c.Mode] {
		return fmt.Errorf("invalid stats mode: %s", c.Mode)
	}
	return nil
}

func (c *StatsConfig) BuildFlags() []string {
	var flags []string
	if c.Mode != "" {
		flags = append(flags, "--mode="+c.Mode)
	}
	for _, h := range c.Host {
		flags = append(flags, "--host="+h)
	}
	for _, p := range c.Path {
		flags = append(flags, "--path="+p)
	}
	for _, t := range c.Tag {
		flags = append(flags, "--tag="+t)
	}
	flags = append(flags, c.SnapshotIDs...)
	return flags
}

var _ ConfigSection = (*StatsConfig)(nil)
