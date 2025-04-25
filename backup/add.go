package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewAddCommand() *cobra.Command {
	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a new worktree from a remote branch",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Add a new worktree\n")
			fmt.Printf("<< TODO >>\n")
		},
	}
	return addCmd
}