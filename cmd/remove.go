package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var taskToRemove string

var RemoveCmd = &cobra.Command{
    Use:   "remove",
    Short: "Remove a task from the to-do list",
    Run: func(cmd *cobra.Command, args []string) {
        for i, task := range Tasks {
            if task == taskToRemove {
                Tasks = append(Tasks[:i], Tasks[i+1:]...)
                fmt.Printf("Removed task: %s\n", taskToRemove)
                return
            }
        }
        fmt.Printf("Task not found: %s\n", taskToRemove)
    },
}

func init() {
    RemoveCmd.Flags().StringVarP(&taskToRemove, "task", "t", "", "Task to be removed")
    RemoveCmd.MarkFlagRequired("task")
}

