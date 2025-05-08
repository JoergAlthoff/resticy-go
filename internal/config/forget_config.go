package config

import "fmt"

type ForgetConfig struct {
	KeepLast          int      `yaml:"keep-last,omitempty"`
	KeepHourly        int      `yaml:"keep-hourly,omitempty"`
	KeepDaily         int      `yaml:"keep-daily,omitempty"`
	KeepWeekly        int      `yaml:"keep-weekly,omitempty"`
	KeepMonthly       int      `yaml:"keep-monthly,omitempty"`
	KeepYearly        int      `yaml:"keep-yearly,omitempty"`
	KeepWithin        string   `yaml:"keep-within,omitempty"`
	KeepWithinHourly  string   `yaml:"keep-within-hourly,omitempty"`
	KeepWithinDaily   string   `yaml:"keep-within-daily,omitempty"`
	KeepWithinWeekly  string   `yaml:"keep-within-weekly,omitempty"`
	KeepWithinMonthly string   `yaml:"keep-within-monthly,omitempty"`
	KeepWithinYearly  string   `yaml:"keep-within-yearly,omitempty"`
	Prune             bool     `yaml:"prune,omitempty"`
	GroupBy           string   `yaml:"group-by,omitempty"`
	DryRun            bool     `yaml:"dry-run,omitempty"`
	Host              []string `yaml:"host,omitempty"`
	Tag               []string `yaml:"tag,omitempty"`
	Path              []string `yaml:"path,omitempty"`
}

func (c *ForgetConfig) ApplyDefaults() {
	if c.KeepLast == 0 {
		c.KeepLast = 5
	}
	if c.GroupBy == "" {
		c.GroupBy = "host"
	}
	if !c.Prune {
		c.Prune = true
	}
}

func (c *ForgetConfig) Validate() error {
	if c.KeepLast == 0 &&
		c.KeepHourly == 0 &&
		c.KeepDaily == 0 &&
		c.KeepWeekly == 0 &&
		c.KeepMonthly == 0 &&
		c.KeepYearly == 0 &&
		c.KeepWithin == "" &&
		c.KeepWithinHourly == "" &&
		c.KeepWithinDaily == "" &&
		c.KeepWithinWeekly == "" &&
		c.KeepWithinMonthly == "" &&
		c.KeepWithinYearly == "" {
		return fmt.Errorf("at least one retention policy (Keep*) must be specified")
	}

	allowedGroupBy := map[string]bool{
		"host":  true,
		"paths": true,
		"tags":  true,
	}
	if c.GroupBy != "" && !allowedGroupBy[c.GroupBy] {
		return fmt.Errorf("invalid group-by value: %s", c.GroupBy)
	}

	return nil
}

var _ ConfigSection = (*ForgetConfig)(nil)