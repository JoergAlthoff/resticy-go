package subcmds

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/JoergAlthoff/resticy-go/internal/config"
)

type StatusCommand struct {
	appConfig *config.AppConfig
}

func (command *StatusCommand) printSnapshots() (string, error) {
	fmt.Println("== Snapshots ==")
	snapshotArgs := append([]string{"snapshots"}, command.appConfig.Parent.BuildFlags()...)
	output, err := runRestic(snapshotArgs, command.appConfig.Debug)
	if err != nil {
		return "", err
	}
	fmt.Println(output)
	return output, nil
}

func (command *StatusCommand) printStats() (string, error) {
	fmt.Println("\n== Stats ==")
	statsArgs := append([]string{"stats"}, command.appConfig.Parent.BuildFlags()...)
	if len(command.appConfig.Stats.SnapshotIDs) == 0 {
		command.appConfig.Stats.SnapshotIDs = []string{"latest"}
	}
	statsArgs = append(statsArgs, command.appConfig.Stats.BuildFlags()...)
	output, err := runRestic(statsArgs, command.appConfig.Debug)
	if err != nil {
		return "", err
	}
	fmt.Println(output)
	return output, nil
}

func (command *StatusCommand) printLocks() error {
	fmt.Println("\n== Locks ==")
	lockArgs := append([]string{"list", "locks"}, command.appConfig.Parent.BuildFlags()...)
	output, err := runRestic(lockArgs, command.appConfig.Debug)
	if err != nil {
		return err
	}

	output = strings.TrimSpace(output)
	if output == "" {
		fmt.Println("✔ No active locks found.")
	} else {
		fmt.Println(output)
		fmt.Println("⚠ Warning: Active repository locks found. Use 'restic unlock' if necessary.")
	}
	return nil
}

func (command *StatusCommand) printSummary(snapshotOutput, statsOutput string) {
	fmt.Println("\n== Summary ==")
	lines := strings.Split(snapshotOutput, "\n")
	var lastTime time.Time
	timePattern := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`)
	for i := len(lines) - 1; i >= 0; i-- {
		if match := timePattern.FindString(lines[i]); match != "" {
			lastTime, _ = time.Parse("2006-01-02 15:04:05", match)
			break
		}
	}
	if !lastTime.IsZero() {
		daysOld := int(time.Since(lastTime).Hours() / 24)
		fmt.Printf("Last Snapshot Time:  %s\n", lastTime.Format(time.RFC1123))
		fmt.Printf("Snapshot Age:        %d day(s)\n", daysOld)
	} else {
		fmt.Println("Last Snapshot Time:  Not found")
	}

	lines = strings.Split(statsOutput, "\n")
	var snapshotsProcessed, fileCount, totalSize string
	for _, line := range lines {
		if strings.Contains(line, "Snapshots processed") {
			snapshotsProcessed = strings.TrimSpace(line)
		}
		if strings.Contains(line, "Total File Count") {
			fileCount = strings.TrimSpace(line)
		}
		if strings.Contains(line, "Total Size") {
			totalSize = strings.TrimSpace(line)
		}
	}
	fmt.Printf("\nStats Summary (mode: %s):\n", command.appConfig.Stats.Mode)
	if snapshotsProcessed != "" {
		fmt.Println("  " + snapshotsProcessed)
	}
	if fileCount != "" {
		fmt.Println("  " + fileCount)
	}
	if totalSize != "" {
		fmt.Println("  " + totalSize)
	}
}

// Execute runs the 'status' subcommand by combining three informative restic commands:
//   1. 'restic snapshots' – shows existing backup snapshots
//   2. 'restic stats' – shows space usage statistics
//   3. 'restic list locks' – checks for active repository locks that may indicate unfinished or failed operations
func (command *StatusCommand) Execute() error {
	snapshotOutput, err := command.printSnapshots()
	if err != nil {
		return err
	}
	statsOutput, err := command.printStats()
	if err != nil {
		return err
	}
	if err := command.printLocks(); err != nil {
		return err
	}
	command.printSummary(snapshotOutput, statsOutput)
	fmt.Println("\n✔ Status completed successfully.")
	return nil
}

func NewStatusCommand(appConfig *config.AppConfig) *StatusCommand {
	return &StatusCommand{appConfig: appConfig}
}

// buildArgs implements SubCommand but is unused for StatusCommand
func (command *StatusCommand) buildArgs() {
	if command.appConfig.Debug {
		fmt.Printf("cfg.Parent content: %+v\n", command.appConfig.Parent)

		fmt.Printf("Built arguments for snapshots: %v\n",
			append([]string{"snapshots"}, command.appConfig.Parent.BuildFlags()...))

		fmt.Printf("Built arguments for stats: %v\n",
			append([]string{"stats"}, command.appConfig.Parent.BuildFlags()...))

		fmt.Printf("Built arguments for locks: %v\n",
			append([]string{"list", "locks"}, command.appConfig.Parent.BuildFlags()...))
	}
}

var _ SubCommand = (*StatusCommand)(nil)
