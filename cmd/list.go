package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
    Use:   "list",
    Short: "List all tasks",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Here are your tasks:")
        for i, task := range Tasks {
            fmt.Printf("%d: %s\n", i+1, task)
        }
    },
}

func init() {
    // Initialization logic for list command
}

