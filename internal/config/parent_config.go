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


func (g *ParentConfig) buildRepositoryFlags() []string {
	var flags []string
	if g.Repository != "" {
		flags = append(flags, "--repo="+g.Repository)
	}
	if g.RepositoryFile != "" {
		flags = append(flags, "--repository-file="+g.RepositoryFile)
	}
	return flags
}

func (g *ParentConfig) buildAuthFlags() []string {
	var flags []string
	if g.PasswordFile != "" {
		flags = append(flags, "--password-file="+g.PasswordFile)
	}
	if g.PasswordCommand != "" {
		flags = append(flags, "--password-command="+g.PasswordCommand)
	}
	if g.KeyHint != "" {
		flags = append(flags, "--key-hint="+g.KeyHint)
	}
	return flags
}

func (g *ParentConfig) buildCacheFlags() []string {
	var flags []string
	if g.CacheDir != "" {
		flags = append(flags, "--cache-dir="+g.CacheDir)
	}
	if g.CleanupCache {
		flags = append(flags, "--cleanup-cache")
	}
	if g.NoCache {
		flags = append(flags, "--no-cache")
	}
	return flags
}

func (g *ParentConfig) buildNetworkFlags() []string {
	var flags []string
	if g.CACert != "" {
		flags = append(flags, "--cacert="+g.CACert)
	}
	if g.InsecureTLS {
		flags = append(flags, "--insecure-tls")
	}
	if g.TLSClientCert != "" {
		flags = append(flags, "--tls-client-cert="+g.TLSClientCert)
	}
	if g.LimitDownload > 0 {
		flags = append(flags, fmt.Sprintf("--limit-download=%d", g.LimitDownload))
	}
	if g.LimitUpload > 0 {
		flags = append(flags, fmt.Sprintf("--limit-upload=%d", g.LimitUpload))
	}
	return flags
}

func (g *ParentConfig) buildOperationFlags() []string {
	var flags []string
	if g.Compression != "" {
		flags = append(flags, "--compression="+g.Compression)
	}
	if g.PackSize > 0 {
		flags = append(flags, fmt.Sprintf("--pack-size=%d", g.PackSize))
	}
	if g.NoLock {
		flags = append(flags, "--no-lock")
	}
	if g.RetryLock != "" {
		flags = append(flags, "--retry-lock="+g.RetryLock)
	}
	if g.NoExtraVerify {
		flags = append(flags, "--no-extra-verify")
	}
	return flags
}

func (g *ParentConfig) buildOutputFlags() []string {
	var flags []string
	if g.JSON {
		flags = append(flags, "--json")
	}
	if g.Quiet {
		flags = append(flags, "--quiet")
	}
	if g.Verbose > 0 {
		flags = append(flags, fmt.Sprintf("--verbose=%d", g.Verbose))
	}
	for _, opt := range g.Option {
		flags = append(flags, "--option="+opt)
	}
	return flags
}

func (g *ParentConfig) BuildFlags() []string {
	flags := make([]string, 0)
	flags = append(flags, g.buildRepositoryFlags()...)
	flags = append(flags, g.buildAuthFlags()...)
	flags = append(flags, g.buildCacheFlags()...)
	flags = append(flags, g.buildNetworkFlags()...)
	flags = append(flags, g.buildOperationFlags()...)
	flags = append(flags, g.buildOutputFlags()...)
	return flags
}

var _ ConfigSection = (*ParentConfig)(nil)