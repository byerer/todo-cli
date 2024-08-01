package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Todo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1.0.0")
	},
}
