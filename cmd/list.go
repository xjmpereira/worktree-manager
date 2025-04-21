package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all worktrees that currently exist.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Listing worktrees\n")
		fmt.Printf("<< TODO >>\n")
	},
}
