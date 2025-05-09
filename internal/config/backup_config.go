package config

import "fmt"

type BackupConfig struct {
	Exclude             []string `yaml:"exclude,omitempty"`
	FilesFrom           string   `yaml:"files-from,omitempty"`
	FilesFromRaw        string   `yaml:"files-from-raw,omitempty"`
	FilesFromVerbatim   []string `yaml:"files-from-verbatim,omitempty"`
	Stdin               string   `yaml:"stdin,omitempty"`
	StdinFilename       string   `yaml:"stdin-filename,omitempty"`
	ExcludeFile         []string `yaml:"exclude-file,omitempty"`
	ExcludeIfPresent    []string `yaml:"exclude-if-present,omitempty"`
	ExcludeLargerThan   string   `yaml:"exclude-larger-than,omitempty"`
	ExcludeCaches       bool     `yaml:"exclude-caches,omitempty"`
	IExclude            []string `yaml:"iexclude,omitempty"`
	IExcludeFile        []string `yaml:"iexclude-file,omitempty"`
	DryRun              bool     `yaml:"dry-run,omitempty"`
	Force               bool     `yaml:"force,omitempty"`
	IgnoreCtime         bool     `yaml:"ignore-ctime,omitempty"`
	IgnoreInode         bool     `yaml:"ignore-inode,omitempty"`
	NoScan              bool     `yaml:"no-scan,omitempty"`
	OneFileSystem       bool     `yaml:"one-file-system,omitempty"`
	WithAtime           bool     `yaml:"with-atime,omitempty"`
	ReadConcurrency     int      `yaml:"read-concurrency,omitempty"`
	Host                string   `yaml:"host,omitempty"`
	Tag                 []string `yaml:"tag,omitempty"`
	GroupBy             string   `yaml:"group-by,omitempty"`
	Parent              string   `yaml:"parent,omitempty"`
	Time                string   `yaml:"time,omitempty"`
	Help                bool     `yaml:"help,omitempty"`
}

func (backupConfig *BackupConfig) BuildFlags() []string {
	var flags []string

	for _, val := range backupConfig.Exclude {
		flags = append(flags, "--exclude="+val)
	}
	for _, val := range backupConfig.ExcludeFile {
		flags = append(flags, "--exclude-file="+val)
	}
	for _, val := range backupConfig.ExcludeIfPresent {
		flags = append(flags, "--exclude-if-present="+val)
	}
	for _, val := range backupConfig.IExclude {
		flags = append(flags, "--iexclude="+val)
	}
	for _, val := range backupConfig.IExcludeFile {
		flags = append(flags, "--iexclude-file="+val)
	}
	for _, val := range backupConfig.FilesFromVerbatim {
		flags = append(flags, "--files-from-verbatim="+val)
	}
	for _, val := range backupConfig.Tag {
		flags = append(flags, "--tag="+val)
	}

	return flags
}

func (backupConfig *BackupConfig) ApplyDefaults() {
	// Beispiel für einen sinnvollen Standardwert
	if backupConfig.ExcludeLargerThan == "" {
		backupConfig.ExcludeLargerThan = "5G"
	}

	// Standardwerte für booleans, falls benötigt
	// (nicht zwingend erforderlich, da bools in Go defaultmäßig false sind)
}

func (backupConfig *BackupConfig) Validate() error {
	if backupConfig.ReadConcurrency < 0 {
		return fmt.Errorf("read-concurrency must not be negative")
	}
	if backupConfig.ExcludeLargerThan != "" && backupConfig.ExcludeLargerThan[0] == '-' {
		return fmt.Errorf("exclude-larger-than must not be negative")
	}
	// Weitere Regeln können bei Bedarf ergänzt werden
	return nil
}

var _ ConfigSection = (*BackupConfig)(nil)