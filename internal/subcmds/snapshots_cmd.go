package subcmds

import (
	"fmt"
	"github.com/JoergAlthoff/resticy-go/internal/config"
	"github.com/JoergAlthoff/resticy-go/internal/logging"
)

// SnapshotsCommand represents the 'snapshots' subcommand.
type SnapshotsCommand struct {
	appConfig *config.AppConfig
	args      []string
}

func NewSnapshotsCommand(appConfig *config.AppConfig) *SnapshotsCommand {
	return &SnapshotsCommand{appConfig: appConfig}
}

func (command *SnapshotsCommand) Execute() error {
	fmt.Println("🗂️ Starting snapshots operation...")
	command.buildArgs()
	output, err := runRestic(command.args, command.appConfig.Debug)
	if err != nil {
		return err
	}
	fmt.Println("✅ Snapshots completed. Logging output...")
	err = logging.LogCommandOutput(command.appConfig.InfoLog, "snapshots", output)
	if err != nil {
		return err
	}
	fmt.Println("📝 Snapshots result logged to:", command.appConfig.InfoLog)
	return nil
}

func (command *SnapshotsCommand) buildArgs() {
	if command.appConfig.Debug {
		fmt.Printf("cfg.Parent content: %+v\n", command.appConfig.Parent)
	}

	command.args = append([]string{"snapshots"}, command.appConfig.Parent.BuildFlags()...)
	command.args = append(command.args, command.appConfig.Snapshots.BuildFlags()...)

	if command.appConfig.Debug {
		fmt.Printf("Built arguments: %v\n", command.args)
	}
}

var _ SubCommand = (*SnapshotsCommand)(nil)