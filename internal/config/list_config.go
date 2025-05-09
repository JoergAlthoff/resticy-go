package config

import "fmt"

// ListConfig holds configuration for the 'restic list' command.
type ListConfig struct {
	Type string `yaml:"type,omitempty"` // e.g., "locks", "snapshots", "blobs", etc.
}

// ApplyDefaults sets default values for ListConfig.
func (c *ListConfig) ApplyDefaults() {
	if c.Type == "" {
		c.Type = "locks"
	}
}

// Validate checks if the configured type is supported.
func (c *ListConfig) Validate() error {
	allowed := map[string]bool{
		"locks":     true,
		"snapshots": true,
		"blobs":     true,
		"packs":     true,
		"index":     true,
		"keys":      true,
	}

	if !allowed[c.Type] {
		return fmt.Errorf("unsupported list type: %s", c.Type)
	}
	return nil
}