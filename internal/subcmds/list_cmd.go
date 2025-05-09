package subcmds

import (
	"fmt"
	"github.com/JoergAlthoff/resticy-go/internal/config"
)

// ListCommand executes 'restic list' with a configurable type (e.g. locks, snapshots).
type ListCommand struct {
	appConfig *config.AppConfig
	args      []string
}

// NewList creates a new instance of ListCommand.
func NewList(appConfig *config.AppConfig) *ListCommand {
	return &ListCommand{
		appConfig: appConfig,
	}
}

// buildArgs assembles the full argument list for 'restic list <type>'.
func (command *ListCommand) buildArgs() {
	if command.appConfig.Debug {
		fmt.Printf("cfg.Parent content: %+v\n", command.appConfig.Parent)
	}

	command.args = append([]string{"list", command.appConfig.List.Type}, command.appConfig.Parent.BuildFlags()...)

	if command.appConfig.Debug {
		fmt.Printf("Built arguments: %v\n", command.args)
	}
}

// Execute runs the 'restic list' command.
func (command *ListCommand) Execute() error {
	command.buildArgs()

	output, err := runRestic(command.args, command.appConfig.Debug)
	if err != nil {
		return err
	}

	fmt.Println(output)
	return nil
}

// Ensure ListCommand implements the SubCommand interface.
var _ SubCommand = (*ListCommand)(nil)