package subcmds

import (
    "fmt"
    "github.com/JoergAlthoff/resticy-go/internal/config"
    "github.com/JoergAlthoff/resticy-go/internal/logging"
)

type PruneCommand struct {
    appConfig *config.AppConfig
    args      []string
}

func NewPrune(appConfig *config.AppConfig) *PruneCommand {
    return &PruneCommand{appConfig: appConfig}
}

func (command *PruneCommand) Execute() error {
    fmt.Println("🧹 Starting prune operation...")
    command.buildArgs()
    output, err := runRestic(command.args, command.appConfig.Debug)
    if err != nil {
        return err
    }
    fmt.Println("✅ Prune completed. Logging output...")
    err = logging.LogCommandOutput(command.appConfig.PruneLog, output)
    if err != nil {
        return err
    }
    fmt.Println("📝 Prune result logged to:", command.appConfig.PruneLog)
    return nil
}

func (command *PruneCommand) buildArgs() {
    if command.appConfig.Debug {
        fmt.Printf("cfg.Parent content: %+v\n", command.appConfig.Parent)
    }

    command.args = append([]string{"prune"}, command.appConfig.Parent.BuildFlags()...)
    command.args = append(command.args, command.appConfig.Prune.BuildFlags()...)

    if command.appConfig.Debug {
        fmt.Printf("Built arguments: %v\n", command.args)
    }
}

var _ SubCommand = (*PruneCommand)(nil)
