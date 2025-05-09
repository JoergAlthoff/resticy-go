package subcmds

// SubCommand is the common interface for all executable resticy subcommands.
type SubCommand interface {
	// Execute runs the configured restic command.
	Execute() error
	// buildArgs constructs the CLI arguments for restic.
	buildArgs()
}
