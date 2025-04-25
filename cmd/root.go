package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var gitwsDir string

func NewRootCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "git-ws",
		Short: "A program to manage git worktrees",
		Long: `GitWS is a CLI tool that allows users to easily manage Git Worktrees`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			updateGitwsRoot()
		},
	}
	rootCmd.PersistentFlags().StringVarP(&gitwsDir, "currentdir", "C", "", "Use this as current directory")
	rootCmd.AddCommand(NewRmCommand())
	rootCmd.AddCommand(NewListCommand())
	rootCmd.AddCommand(NewCreateCommand())
	rootCmd.AddCommand(NewConfigCommand())
	rootCmd.AddCommand(NewCloneCommand())
	rootCmd.AddCommand(NewAddCommand())
	rootCmd.AddCommand(NewVersionCommand())

	return rootCmd
}

func Execute() {
	err := NewRootCommand().Execute()
	if err != nil {
		os.Exit(1)
	}
}
