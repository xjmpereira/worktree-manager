package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new worktree from a branch name",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Create a new worktree\n")
		fmt.Printf("<< TODO >>\n")
	},
}
