package subcmds

import (
	"fmt"

	"github.com/JoergAlthoff/resticy-go/internal/config"
	"github.com/JoergAlthoff/resticy-go/internal/logging"
)

// StatsCommand represents the 'stats' subcommand
type StatsCommand struct {
	appConfig *config.AppConfig
	args      []string
}

func NewStatsCommand(appConfig *config.AppConfig) *StatsCommand {
	return &StatsCommand{
		appConfig: appConfig,
		args:      []string{},
	}
}

func (command *StatsCommand) Execute() error {
	fmt.Println("📊 Starting stats operation...")
	command.buildArgs()
	output, err := runRestic(command.args, command.appConfig.Debug)
	if err != nil {
		return err
	}
	fmt.Println("✅ Stats completed. Logging output...")
	err = logging.LogCommandOutput(command.appConfig.InfoLog, "stats", output)
	if err != nil {
		return err
	}
	fmt.Println("📝 Stats result logged to:", command.appConfig.InfoLog)
	return nil
}


func (command *StatsCommand) buildArgs() {
	if command.appConfig.Debug {
		fmt.Printf("cfg.Parent content: %+v\n", command.appConfig.Parent)
	}

	command.args = append([]string{"stats"}, command.appConfig.Parent.BuildFlags()...)
	command.args = append(command.args, command.appConfig.Stats.BuildFlags()...)

	if command.appConfig.Debug {
		fmt.Printf("Built arguments: %v\n", command.args)
	}
}

var _ SubCommand = (*StatsCommand)(nil)