package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"resticy-go/internal/config"
	"resticy-go/internal/subcmds"
)

func main() {
	flagDebug := flag.Bool("debug", false, "Enable debug output from resticy")
	flagConfig := flag.String("config", "", "Path to configuration YAML file")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 || *flagConfig == "" {
		fmt.Fprintln(os.Stderr, "Usage: resticy <command> --config <file.yaml> [--debug]")
		os.Exit(1)
	}

	command := args[0]
	appCconfigPath := *flagConfig

	cfg, err := config.Load(appCconfigPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Configuration error: %v\n", err)
		os.Exit(1)
	}

	if *flagDebug {
		fmt.Fprintln(os.Stderr, "Debug mode active")
	}

	// Validate configuration based on specific subcommand requirements
	if err := validateConfigForCommand(command, cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Invalid config for '%s': %v\n", command, err)
		os.Exit(1)
	}

	switch command {
	case "backup":
		if *flagDebug {
			fmt.Println("Executing backup...")
		}
		// TODO: Implement backup logic
	case "forget":
		if *flagDebug {
			fmt.Println("Executing forget...")
		}
		// TODO: Implement forget logic
		
	case "check":
		if *flagDebug {
			fmt.Println("Executing check...")
		}
		if err := subcmds.NewCheck(cfg).Execute(); err != nil {
			fmt.Fprintf(os.Stderr, "restic check failed: %v\n", err)
			os.Exit(1)
		}

	case "snapshots":
		if *flagDebug {
			fmt.Println("Executing snapshots...")
		}
		// TODO: Implement snapshots logic
	case "diff":
		if *flagDebug {
			fmt.Println("Executing diff...")
		}
		// TODO: Implement diff logic
	case "prune":
		if *flagDebug {
			fmt.Println("Executing prune...")
		}
		// TODO: Implement prune logic
	case "init":
		if *flagDebug {
			fmt.Println("Executing init...")
		}
		// TODO: Implement init logic
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}
}

func validateConfigForCommand(cmd string, cfg *config.AppConfig) error {
	switch cmd {
	case "backup":
		if cfg.Parent.Repository == "" || cfg.Parent.RepositoryFile == "" {
			return errors.New("missing required fields: repository or password_file")
		}
	case "forget":
		if cfg.Parent.RepositoryFile == "" {
			return errors.New("missing required field: repository")
		}
	case "check":
		// validation logic for "check" can be added here
	case "snapshots":
		// no required validation yet
	case "diff":
		// no required validation yet
	case "prune":
		if cfg.Parent.RepositoryFile == "" {
			return errors.New("missing required field: repository")
		}
	case "init":
		// no required validation yet
	}
	return nil
}
