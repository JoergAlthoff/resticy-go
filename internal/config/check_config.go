package config

type CheckConfig struct {
	ReadData    bool  `yaml:"read_data,omitempty"`
	CheckUnused bool  `yaml:"check_unused,omitempty"`
	ReadDataSubset string `yaml:"read_data_subset,omitempty"`
	WithCache   *bool `yaml:"with_cache,omitempty"`
}

func (c *CheckConfig) ApplyDefaults() {
	if c.WithCache == nil {
		defaultVal := true
		c.WithCache = &defaultVal
	}
}

func (c *CheckConfig) Validate() error {
	return nil
}
func (c *CheckConfig) BuildFlags() []string {
	var flags []string

	if c.ReadData {
		flags = append(flags, "--read-data")
	}
	if c.ReadDataSubset != "" {
		flags = append(flags, "--read-data-subset="+c.ReadDataSubset)
	}
	if c.WithCache != nil && *c.WithCache {
		flags = append(flags, "--with-cache")
	}
	if c.CheckUnused {
		flags = append(flags, "--check-unused")
	}

	return flags
}

var _ ConfigSection = (*CheckConfig)(nil)