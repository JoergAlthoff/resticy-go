package config

type CheckConfig struct {
	ReadData    bool  `yaml:"read_data,omitempty"`
	CheckUnused bool  `yaml:"check_unused,omitempty"`
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
