package subcmds

import (
	"fmt"
	"github.com/JoergAlthoff/resticy-go/internal/config"
	"github.com/JoergAlthoff/resticy-go/internal/logging"
)

// ForgetCommand handles the 'forget' subcommand.
type ForgetCommand struct {
	appConfig *config.AppConfig
	args      []string
}

// NewForgetCommand constructs a new ForgetCommand instance.
func NewForgetCommand(appConfig *config.AppConfig) *ForgetCommand {
	return &ForgetCommand{
		appConfig: appConfig,
	}
}

func (command *ForgetCommand) buildArgs() {
	if command.appConfig.Debug {
		fmt.Printf("cfg.Parent content: %+v\n", command.appConfig.Parent)
	}

	command.args = append([]string{"forget"}, command.appConfig.Parent.BuildFlags()...)
	command.args = append(command.args, command.appConfig.Forget.BuildFlags()...)

	if command.appConfig.Debug {
		fmt.Printf("Built arguments: %v\n", command.args)
	}
}

func (command *ForgetCommand) Execute() error {
	command.buildArgs()
	output, err := runRestic(command.args, command.appConfig.Debug)
	if err != nil {
		return err
	}
	return logging.LogCommandOutput(command.appConfig.ForgetLog, output)
}

// Interface check
var _ SubCommand = (*ForgetCommand)(nil)
