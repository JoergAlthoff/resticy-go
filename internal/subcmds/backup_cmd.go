package subcmds

import (
	"fmt"
	"os"
	"github.com/JoergAlthoff/resticy-go/internal/config"
	"github.com/JoergAlthoff/resticy-go/internal/logging"
)

type BackupCommand struct {
	appConfig     *config.AppConfig
	args    []string
	sources []string
}

func NewBackup(appConfig *config.AppConfig, sources []string) *BackupCommand {
	return &BackupCommand{appConfig: appConfig, sources: sources}
}

func (command *BackupCommand) buildArgs() {
	if command.appConfig.Debug {
		fmt.Printf("cfg.Parent content: %+v\n", command.appConfig.Parent)
	}

	command.args = append([]string{"backup"}, command.appConfig.Parent.BuildFlags()...)
	command.args = append(command.args, command.sources...)
	command.args = append(command.args, command.appConfig.Backup.BuildFlags()...)

	if command.appConfig.Debug {
		fmt.Printf("Built arguments: %v\n", command.args)
	}
}

func (command *BackupCommand) Execute() error {
	if len(command.sources) == 0 {
		fmt.Fprintln(os.Stderr, "Error: No backup source path provided.")
		os.Exit(1)
	}

	command.buildArgs()
	output, err := runRestic(command.args, command.appConfig.Debug)
	if err != nil {
		return err
	}
	return logging.LogCommandOutput(command.appConfig.ForgetLog, output)
}


// Interface check
var _ SubCommand = (*BackupCommand)(nil)