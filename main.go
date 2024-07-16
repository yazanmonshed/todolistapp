package main

import (
    "todolist/cmd"
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "todolist",
    Short: "ToDo List is a CLI application to manage your tasks.",
    Long:  "A simple To-Do list application written in Go using Cobra library.",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Welcome to ToDo List CLI application. Use -h to see available commands.")
    },
}

func main() {
    rootCmd.AddCommand(cmd.AddCmd)
    rootCmd.AddCommand(cmd.ListCmd)
    rootCmd.AddCommand(cmd.RemoveCmd)

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

