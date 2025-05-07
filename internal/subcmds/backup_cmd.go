package subcmds

import (
	"fmt"
	"os"
	"resticy-go/internal/config"
)

type Backup struct {
	cfg     *config.AppConfig
	args    []string
	sources []string
}

func NewBackup(cfg *config.AppConfig, sources []string) *Backup {
	return &Backup{cfg: cfg, sources: sources}
}

func (b *Backup) buildArgs() {
	b.args = []string{"backup"}

	b.args = append(b.args, b.sources...)

	if b.cfg.Parent.Repository != "" {
		b.args = append(b.args, "--repo", b.cfg.Parent.Repository)
	}

	if b.cfg.Parent.PasswordFile != "" {
		b.args = append(b.args, "--password-file", b.cfg.Parent.PasswordFile)
	}

	b.args = append(b.args, b.cfg.BuildExcludeFlags()...)

	if b.cfg.Debug {
		fmt.Printf("Built arguments: %v\n", b.args)
	}
}

func (b *Backup) Execute() error {
	if len(b.sources) == 0 {
		fmt.Fprintln(os.Stderr, "Error: No backup source path provided.")
		os.Exit(1)
	}

	b.buildArgs()
	return runRestic(b.args, b.cfg.Parent.Verbose)
}