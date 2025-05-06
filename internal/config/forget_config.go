package config

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
