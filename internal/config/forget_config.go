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

func (forgetConfig *ForgetConfig) BuildFlags() []string {
	var flags []string
	flags = append(flags, forgetConfig.snapshotRetentionFlags()...)
	flags = append(flags, forgetConfig.timeBasedRetentionFlags()...)
	flags = append(flags, forgetConfig.cleanupBehaviorFlags()...)
	return flags
}

func (forgetConfig *ForgetConfig) snapshotRetentionFlags() []string {
	var flags []string

	if forgetConfig.KeepLast > 0 {
		flags = append(flags, fmt.Sprintf("--keep-last=%d", forgetConfig.KeepLast))
	}
	if forgetConfig.KeepHourly > 0 {
		flags = append(flags, fmt.Sprintf("--keep-hourly=%d", forgetConfig.KeepHourly))
	}
	if forgetConfig.KeepDaily > 0 {
		flags = append(flags, fmt.Sprintf("--keep-daily=%d", forgetConfig.KeepDaily))
	}
	if forgetConfig.KeepWeekly > 0 {
		flags = append(flags, fmt.Sprintf("--keep-weekly=%d", forgetConfig.KeepWeekly))
	}
	if forgetConfig.KeepMonthly > 0 {
		flags = append(flags, fmt.Sprintf("--keep-monthly=%d", forgetConfig.KeepMonthly))
	}
	if forgetConfig.KeepYearly > 0 {
		flags = append(flags, fmt.Sprintf("--keep-yearly=%d", forgetConfig.KeepYearly))
	}

	return flags
}

func (forgetConfig *ForgetConfig) timeBasedRetentionFlags() []string {
	var flags []string

	if forgetConfig.KeepWithin != "" {
		flags = append(flags, "--keep-within="+forgetConfig.KeepWithin)
	}
	if forgetConfig.KeepWithinHourly != "" {
		flags = append(flags, "--keep-within-hourly="+forgetConfig.KeepWithinHourly)
	}
	if forgetConfig.KeepWithinDaily != "" {
		flags = append(flags, "--keep-within-daily="+forgetConfig.KeepWithinDaily)
	}
	if forgetConfig.KeepWithinWeekly != "" {
		flags = append(flags, "--keep-within-weekly="+forgetConfig.KeepWithinWeekly)
	}
	if forgetConfig.KeepWithinMonthly != "" {
		flags = append(flags, "--keep-within-monthly="+forgetConfig.KeepWithinMonthly)
	}
	if forgetConfig.KeepWithinYearly != "" {
		flags = append(flags, "--keep-within-yearly="+forgetConfig.KeepWithinYearly)
	}

	return flags
}

func (forgetConfig *ForgetConfig) cleanupBehaviorFlags() []string {
	var flags []string

	if forgetConfig.Prune {
		flags = append(flags, "--prune")
	}
	
	if forgetConfig.DryRun {
		flags = append(flags, "--dry-run")
	}

	if forgetConfig.GroupBy != "" {
		flags = append(flags, "--group-by="+forgetConfig.GroupBy)
	}

	for _, host := range forgetConfig.Host {
		flags = append(flags, "--host="+host)
	}

	for _, tag := range forgetConfig.Tag {
		flags = append(flags, "--tag="+tag)
	}

	for _, path := range forgetConfig.Path {
		flags = append(flags, "--path="+path)
	}

	return flags
}

var _ ConfigSection = (*ForgetConfig)(nil)