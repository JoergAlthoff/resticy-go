package config

// ConfigSection is the common interface for all configuration blocks.
type ConfigSection interface {
	ApplyDefaults()
	Validate() error
	BuildFlags() []string
}