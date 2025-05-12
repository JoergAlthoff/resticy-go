package subcmds

import (
	"fmt"
	"os"
	"time"

	"github.com/JoergAlthoff/resticy-go/internal/config"
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
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	output = fmt.Sprintf("[%s] restic backup started\n%s", timestamp, output)
	logFile := command.appConfig.BackupLog
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file %s: %w", logFile, err)
	}
	defer file.Close()

	_, err = file.WriteString(output + "\n")
	if err != nil {
		return fmt.Errorf("failed to write to log file %s: %w", logFile, err)
	}
	return nil
}


// Interface check
var _ SubCommand = (*BackupCommand)(nil)