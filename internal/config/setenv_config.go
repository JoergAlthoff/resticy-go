package config

type SetenvConfig struct {
	OutputFile string `yaml:"output_file"`
	Debug      bool   `yaml:"debug"`
}