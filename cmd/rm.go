package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rmCmd)
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a worktree and clean it up",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Remove a worktree\n")
		fmt.Printf("<< TODO >>\n")
	},
}
