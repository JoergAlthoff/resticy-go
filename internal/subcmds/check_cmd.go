package subcmds

import (
	"fmt"
	"os"
	"os/exec"
	"resticy-go/internal/config"
)

type Check struct {
	cfg  *config.AppConfig
	args []string
}

func NewCheck(cfg *config.AppConfig) *Check {
	return &Check{cfg: cfg}
}

func (c *Check) Execute() error {
	c.buildArgs()

	cmd := exec.Command("restic", c.args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if c.cfg.Parent.Verbose > 0 {
		fmt.Printf("Executing: restic %v\n", c.args)
	}

	return cmd.Run()
}

func (c *Check) buildArgs() {
	c.args = []string{"check"}

	if c.cfg.Parent.RepositoryFile != "" {
		c.args = append(c.args, "--repository-file", c.cfg.Parent.RepositoryFile)
	} else if c.cfg.Parent.Repository != "" {
		c.args = append(c.args, "--repo", c.cfg.Parent.Repository)
	}

	if c.cfg.Parent.PasswordCommand != "" {
		c.args = append(c.args, "--password-command", c.cfg.Parent.PasswordCommand)
	} else if c.cfg.Parent.PasswordFile != "" {
		c.args = append(c.args, "--password-file", c.cfg.Parent.PasswordFile)
	}

	if c.cfg.Parent.CACert != "" {
		c.args = append(c.args, "--cacert", c.cfg.Parent.CACert)
	}

	if c.cfg.Parent.Quiet {
		c.args = append(c.args, "--quiet")
	}

	if c.cfg.Parent.Verbose > 0 {
		c.args = append(c.args, fmt.Sprintf("--verbose=%d", c.cfg.Parent.Verbose))
	}
}
