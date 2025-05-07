package config

import (
	"fmt"
	"os"
)

type ParentConfig struct {
	Repository      string   `yaml:"repository,omitempty"`
	RepositoryFile  string   `yaml:"repository_file,omitempty"`
	PasswordFile    string   `yaml:"password_file,omitempty"`
	PasswordCommand string   `yaml:"password_command,omitempty"`
	KeyHint         string   `yaml:"key_hint,omitempty"`
	CacheDir        string   `yaml:"cache_dir,omitempty"`
	CleanupCache    bool     `yaml:"cleanup_cache,omitempty"`
	NoCache         bool     `yaml:"no_cache,omitempty"`
	CACert          string   `yaml:"cacert,omitempty"`
	InsecureTLS     bool     `yaml:"insecure_tls,omitempty"`
	Compression     string   `yaml:"compression,omitempty"`
	PackSize        int      `yaml:"pack_size,omitempty"`
	NoLock          bool     `yaml:"no_lock,omitempty"`
	RetryLock       string   `yaml:"retry_lock,omitempty"`
	TLSClientCert   string   `yaml:"tls_client_cert,omitempty"`
	LimitDownload   int      `yaml:"limit_download,omitempty"`
	LimitUpload     int      `yaml:"limit_upload,omitempty"`
	NoExtraVerify   bool     `yaml:"no_extra_verify,omitempty"`
	Option          []string `yaml:"option,omitempty"`
	JSON            bool     `yaml:"json,omitempty"`
	Quiet           bool     `yaml:"quiet,omitempty"`
	Verbose         int      `yaml:"verbose,omitempty"`
}

var _ ConfigSection = (*ParentConfig)(nil)

func (g *ParentConfig) ApplyDefaults() {
	if g.Compression == "" {
		g.Compression = os.Getenv("RESTIC_COMPRESSION")
		if g.Compression == "" {
			g.Compression = "auto"
		}
	}

	if g.CACert == "" {
		g.CACert = os.Getenv("RESTIC_CACERT")
	}

	if g.TLSClientCert == "" {
		g.TLSClientCert = os.Getenv("RESTIC_TLS_CLIENT_CERT")
	}

	if g.KeyHint == "" {
		g.KeyHint = os.Getenv("RESTIC_KEY_HINT")
	}

	if g.PackSize == 0 {
		if val := os.Getenv("RESTIC_PACK_SIZE"); val != "" {
			fmt.Sscanf(val, "%d", &g.PackSize)
		}
	}

	if g.PasswordCommand == "" {
		g.PasswordCommand = os.Getenv("RESTIC_PASSWORD_COMMAND")
	}

	if g.PasswordFile == "" {
		g.PasswordFile = os.Getenv("RESTIC_PASSWORD_FILE")
	}

	if g.Repository == "" {
		g.Repository = os.Getenv("RESTIC_REPOSITORY")
	}

	if g.RepositoryFile == "" {
		g.RepositoryFile = os.Getenv("RESTIC_REPOSITORY_FILE")
	}
}

func (g *ParentConfig) Validate() error {

	if g.Repository == "" && g.RepositoryFile == "" {
		return fmt.Errorf("no repository configured (global.repository or global.repository-file)")
	}

	if g.PasswordCommand == "" && g.PasswordFile == "" {
		return fmt.Errorf("no password method configured (global.password-file or global.password-command)")
	}

	if g.Verbose < 0 || g.Verbose > 2 {
		return fmt.Errorf("verbose must be between 0 and 2")
	}

	if g.CACert != "" {
		if _, err := os.Stat(g.CACert); err != nil {
			return fmt.Errorf("cacert file not found: %s", g.CACert)
		}
	}

	if g.PackSize < 0 {
		return fmt.Errorf("pack-size must not be negative")
	}

	return nil
}
