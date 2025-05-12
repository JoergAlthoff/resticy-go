package subcmds

import (
	"fmt"
	"os"
	"strings"

	"github.com/JoergAlthoff/resticy-go/internal/config"
)

type SetenvCommand struct {
	appConfig    *config.AppConfig
	setenvConfig *config.SetenvConfig
	args         []string
}

func NewSetenvCommand(appConfig *config.AppConfig, setenvConfig *config.SetenvConfig) *SetenvCommand {
	return &SetenvCommand{
		appConfig:    appConfig,
		setenvConfig: setenvConfig,
		args:         []string{},
	}
}

func (command *SetenvCommand) SetOutputFile(path string) {
	if path != "" {
		command.setenvConfig.OutputFile = path
	}
}

func (command *SetenvCommand) SetDebug(debug bool) {
	command.setenvConfig.Debug = debug
}

func (command *SetenvCommand) Execute() error {
	cfg := command.appConfig
	if cfg == nil {
		fmt.Fprintln(os.Stderr, "Missing configuration")
		return fmt.Errorf("no configuration")
	}

	script := command.buildExportScript(cfg)
	err := os.WriteFile(command.setenvConfig.OutputFile, []byte(script), 0644)
	if err != nil {
		return fmt.Errorf("failed to write script to %s: %w", command.setenvConfig.OutputFile, err)
	}

	if command.setenvConfig.Debug {
		fmt.Println("Generated script:")
		fmt.Println(script)
	}

	// Hinweis zur Verwendung von 'source' und erklärender Kommentar:
	fmt.Printf("\nTo apply the environment variables in your current shell, run:\n")
	fmt.Printf("  source %s\n", command.setenvConfig.OutputFile)
	fmt.Printf("\nThis will apply only the values defined in your config file.\n")
	fmt.Printf("Existing RESTIC_* environment variables remain unchanged unless explicitly overwritten.\n")

	return nil
}

func (command *SetenvCommand) buildArgs() {
	// intentionally empty; no arguments needed for setenv
}

func (command *SetenvCommand) buildExportScript(cfg *config.AppConfig) string {
	var builder strings.Builder
	builder.WriteString("#!/bin/sh\n")

	command.buildRepositoryScript(cfg, &builder)
	command.buildTLSConfigScript(cfg, &builder)
	command.buildCacheScript(cfg, &builder)
	command.buildPerformanceScript(cfg, &builder)
	command.buildBehaviorScript(cfg, &builder)

	return builder.String()
}

// Utility functions for printing export lines.
func printExportLine(debug bool, builder *strings.Builder, key, value string) {
	if debug {
		fmt.Printf("export %s=%s\n", key, value)
	}
	fmt.Fprintf(builder, "export %s=%s\n", key, value)
}

func printExportBool(debug bool, builder *strings.Builder, key string) {
	if debug {
		fmt.Printf("export %s=1\n", key)
	}
	fmt.Fprintf(builder, "export %s=1\n", key)
}

func (command *SetenvCommand) buildRepositoryScript(cfg *config.AppConfig, builder *strings.Builder) {
	if v := cfg.Parent.Repository; v != "" {
		printExportLine(command.setenvConfig.Debug, builder, "RESTIC_REPOSITORY", v)
	}
	if v := cfg.Parent.RepositoryFile; v != "" {
		printExportLine(command.setenvConfig.Debug, builder, "RESTIC_REPOSITORY_FILE", v)
	}
	if v := cfg.Parent.PasswordFile; v != "" {
		printExportLine(command.setenvConfig.Debug, builder, "RESTIC_PASSWORD_FILE", v)
	}
	if v := cfg.Parent.PasswordCommand; v != "" {
		printExportLine(command.setenvConfig.Debug, builder, "RESTIC_PASSWORD_COMMAND", v)
	}
	if v := cfg.Parent.KeyHint; v != "" {
		printExportLine(command.setenvConfig.Debug, builder, "RESTIC_KEY_HINT", v)
	}
}

func (command *SetenvCommand) buildTLSConfigScript(cfg *config.AppConfig, builder *strings.Builder) {
	if v := cfg.Parent.CACert; v != "" {
		printExportLine(command.setenvConfig.Debug, builder, "RESTIC_CACERT", v)
	}
	if v := cfg.Parent.TLSClientCert; v != "" {
		printExportLine(command.setenvConfig.Debug, builder, "RESTIC_TLS_CLIENT_CERT", v)
	}
	if cfg.Parent.InsecureTLS {
		printExportBool(command.setenvConfig.Debug, builder, "RESTIC_INSECURE_TLS")
	}
}

func (command *SetenvCommand) buildCacheScript(cfg *config.AppConfig, builder *strings.Builder) {
	if v := cfg.Parent.CacheDir; v != "" {
		printExportLine(command.setenvConfig.Debug, builder, "RESTIC_CACHE_DIR", v)
	}
	if cfg.Parent.CleanupCache {
		printExportBool(command.setenvConfig.Debug, builder, "RESTIC_CLEANUP_CACHE")
	}
	if cfg.Parent.NoCache {
		printExportBool(command.setenvConfig.Debug, builder, "RESTIC_NO_CACHE")
	}
}

func (command *SetenvCommand) buildPerformanceScript(cfg *config.AppConfig, builder *strings.Builder) {
	if v := cfg.Parent.Compression; v != "" {
		printExportLine(command.setenvConfig.Debug, builder, "RESTIC_COMPRESSION", v)
	}
	if v := cfg.Parent.PackSize; v > 0 {
		printExportLine(command.setenvConfig.Debug, builder, "RESTIC_PACK_SIZE", fmt.Sprintf("%d", v))
	}
	if v := cfg.Parent.LimitUpload; v > 0 {
		printExportLine(command.setenvConfig.Debug, builder, "RESTIC_LIMIT_UPLOAD", fmt.Sprintf("%d", v))
	}
	if v := cfg.Parent.LimitDownload; v > 0 {
		printExportLine(command.setenvConfig.Debug, builder, "RESTIC_LIMIT_DOWNLOAD", fmt.Sprintf("%d", v))
	}
}

func (command *SetenvCommand) buildBehaviorScript(cfg *config.AppConfig, builder *strings.Builder) {
	if v := cfg.Parent.Verbose; v > 0 {
		printExportLine(command.setenvConfig.Debug, builder, "RESTIC_VERBOSITY", fmt.Sprintf("%d", v))
	}
	if cfg.Parent.NoLock {
		printExportBool(command.setenvConfig.Debug, builder, "RESTIC_NO_LOCK")
	}
}

var _ SubCommand = (*SetenvCommand)(nil)
