package logging

import (
    "fmt"
    "os"
    "time"
)

// LogCommandOutput appends the command output to a log file with a timestamp
func LogCommandOutput(logPath string, output string) error {
    if logPath == "" {
        return nil
    }

    file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    timestamp := time.Now().Format("2006-01-02 15:04:05")
    header := fmt.Sprintf("\n[%s] restic output:\n", timestamp)

    _, err = file.Write([]byte(header))
    if err != nil {
        return err
    }
    _, err = file.Write([]byte(output))
    return err
}