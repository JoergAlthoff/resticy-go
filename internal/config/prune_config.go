package config

import (
	"fmt"
	"strings"
)

func hasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

type PruneConfig struct {
	Enabled       bool   `yaml:"enabled"`
	DryRun        bool   `yaml:"dry-run"`
	MaxUnused     string `yaml:"max-unused"`      // e.g. "5%"
	MaxRepackSize string `yaml:"max-repack-size"` // e.g. "500MiB"
}

// BuildFlags implements ConfigSection.
func (c *PruneConfig) BuildFlags() []string {
	var flags []string
	if c.DryRun {
		flags = append(flags, "--dry-run")
	}
	if c.MaxUnused != "" {
		flags = append(flags, "--max-unused", c.MaxUnused)
	}
	if c.MaxRepackSize != "" {
		flags = append(flags, "--max-repack-size", c.MaxRepackSize)
	}
	return flags
}

func (c *PruneConfig) ApplyDefaults() {
	if c.MaxUnused == "" {
		c.MaxUnused = "5%"
	}
	if c.MaxRepackSize == "" {
		c.MaxRepackSize = "500M"
	}
}

func (c *PruneConfig) Validate() error {
	// Very simple validation — could be extended with regex etc.
	if c.MaxUnused != "" && c.MaxUnused[len(c.MaxUnused)-1] != '%' {
		return fmt.Errorf("max-unused should be a percentage (e.g. \"5%%\")")
	}
	if c.MaxRepackSize != "" && !(hasSuffix(c.MaxRepackSize, "M") || hasSuffix(c.MaxRepackSize, "K") || hasSuffix(c.MaxRepackSize, "G")) {
		return fmt.Errorf("max-repack-size must end with K, M or G (e.g. '500M')")
	}
	return nil
}

var _ ConfigSection = (*PruneConfig)(nil)
