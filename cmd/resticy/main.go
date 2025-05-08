package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"resticy-go/internal/config"
	"resticy-go/internal/subcmds"
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
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}