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

	rootCmd.AddCommand(&cobra.Command{
		Use:   "stats",
		Short: "Show repository statistics",
		RunE: func(cmd *cobra.Command, args []string) error {
			return subcmds.NewStatsCommand(cfg).Execute()
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "snapshots",
		Short: "List snapshots in the repository",
		RunE: func(cmd *cobra.Command, args []string) error {
			return subcmds.NewSnapshotsCommand(cfg).Execute()
		},
	})
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}