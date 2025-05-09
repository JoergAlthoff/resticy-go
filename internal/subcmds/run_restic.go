package subcmds

import (
	"fmt"
	"os/exec"
)

// runRestic executes a restic command with the given arguments.
// If verbose == true, it prints the command before execution.
func runRestic(args []string, verbose bool) (string, error) {
	if verbose {
		fmt.Printf("Executing: restic %v\n", args)
	}

	cmd := exec.Command("restic", args...)
	output, err := cmd.CombinedOutput()

	if verbose {
		fmt.Println(string(output))
	}

	return string(output), err
}
