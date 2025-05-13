package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/JoergAlthoff/resticy-go/internal/config"
	"github.com/JoergAlthoff/resticy-go/internal/subcmds"
)

var (
	cfgPath string
	debug   bool
	cfg     *config.AppConfig
)

var rootCmd = &cobra.Command{
	Use:   "resticy",
	Short: "resticy is a wrapper for restic with structured config support",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		cfg, err = config.Load(cfgPath)
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}
		cfg.Debug = debug
		cfg.ApplyDefaults()
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgPath, "config", "", "Path to configuration YAML file")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug output")
	rootCmd.MarkPersistentFlagRequired("config")

	rootCmd.AddCommand(&cobra.Command{
		Use:   "check",
		Short: "Run restic check",
		RunE: func(cmd *cobra.Command, args []string) error {
			return subcmds.NewCheck(cfg).Execute()
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "backup",
		Short: "Run restic backup",
		RunE: func(cmd *cobra.Command, args []string) error {
			return subcmds.NewBackup(cfg, args).Execute()
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "forget",
		Short: "Run restic forget",
		RunE: func(cmd *cobra.Command, args []string) error {
			return subcmds.NewForgetCommand(cfg).Execute()
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "status",
		Short: "Show repository status including snapshots, stats and locks",
		RunE: func(cmd *cobra.Command, args []string) error {
			return subcmds.NewStatusCommand(cfg).Execute()
		},
	})

	var outputFile string

	setenvCmd := &cobra.Command{
		Use:   "setenv",
		Short: "Print environment variable exports for restic",
		Long: `Generates and executes a shell script that exports environment variables
required by restic based on the loaded configuration.

Options:
  --output=FILE  Name of the shell script to generate (default: set_restic_env.sh)
  --debug        Print the generated script to stdout before execution

Note:
  Only values explicitly set in the configuration are exported.
  Existing environment variables remain unchanged unless overwritten by the config.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			setenvCfg := &config.SetenvConfig{
				OutputFile: outputFile,
				Debug:      debug,
			}
			return subcmds.NewSetenvCommand(cfg, setenvCfg).Execute()
		},
	}

	setenvCmd.Flags().StringVar(&outputFile, "output", "set_restic_env.sh", "Name of the shell script to generate")
	rootCmd.AddCommand(setenvCmd)

	rootCmd.AddCommand(&cobra.Command{
		Use:   "snapshots",
		Short: "List snapshots in the repository",
		RunE: func(cmd *cobra.Command, args []string) error {
			return subcmds.NewSnapshotsCommand(cfg).Execute()
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "prune",
		Short: "Run restic prune",
		RunE: func(cmd *cobra.Command, args []string) error {
			return subcmds.NewPrune(cfg).Execute()
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "stats",
		Short: "Show restic repository statistics",
		RunE: func(cmd *cobra.Command, args []string) error {
			return subcmds.NewStatsCommand(cfg).Execute()
		},
	})
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}