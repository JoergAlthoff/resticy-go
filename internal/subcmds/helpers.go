package subcmds

import (
    "fmt"
    "os"
    "os/exec"
)

// runRestic executes a restic command with the given arguments.
// If verbose > 0, it prints the command before execution.
func runRestic(args []string, verbose int) error {
    if verbose > 0 {
        fmt.Printf("Executing: restic %v\n", args)
    }

    cmd := exec.Command("restic", args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    return cmd.Run()
}
