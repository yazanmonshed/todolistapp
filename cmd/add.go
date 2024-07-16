package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var task string

var AddCmd = &cobra.Command{
    Use:   "add",
    Short: "Add a new task to the to-do list",
    Run: func(cmd *cobra.Command, args []string) {
        Tasks = append(Tasks, task)
        fmt.Printf("Added task: %s\n", task)
    },
}

func init() {
    AddCmd.Flags().StringVarP(&task, "task", "t", "", "Task to be added")
    AddCmd.MarkFlagRequired("task")
}

