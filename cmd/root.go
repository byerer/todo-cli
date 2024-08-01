package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A command-line todo application for managing your tasks",
	Long: `Todo is a command-line task management application that helps you keep track of your daily tasks.
With Todo, you can add, remove, and view tasks in a simple and efficient way. It's designed for those who need a quick and minimal setup to start managing their tasks immediately.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
