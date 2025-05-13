package subcmds

import (
	"fmt"
	"github.com/JoergAlthoff/resticy-go/internal/config"
	"github.com/JoergAlthoff/resticy-go/internal/logging"
)

type CheckCommand struct {
	appConfig *config.AppConfig
	args      []string
}

func NewCheck(appConfig *config.AppConfig) *CheckCommand {
	return &CheckCommand{appConfig: appConfig}
}

func (command *CheckCommand) Execute() error {
	fmt.Println("🔍 Starting check operation...")
	command.buildArgs()
	output, err := runRestic(command.args, command.appConfig.Debug)
	if err != nil {
		return err
	}
	fmt.Println("✅ Check completed. Logging output...")
	err = logging.LogCommandOutput(command.appConfig.InfoLog, "check", output)
	if err != nil {
		return err
	}
	fmt.Println("📝 Check result logged to:", command.appConfig.ForgetLog)
	return nil
}

func (command *CheckCommand) buildArgs() {
	if command.appConfig.Debug {
		fmt.Printf("cfg.Parent content: %+v\n", command.appConfig.Parent)
	}

	command.args = append([]string{"check"}, command.appConfig.Parent.BuildFlags()...)
	command.args = append(command.args, command.appConfig.Check.BuildFlags()...)

	if command.appConfig.Debug {
		fmt.Printf("Built arguments: %v\n", command.args)
	}
}

var _ SubCommand = (*CheckCommand)(nil)
