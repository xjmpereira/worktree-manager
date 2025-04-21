package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var gitwsDir string

func init() {
	rootCmd.PersistentFlags().StringVarP(&gitwsDir, "currentdir", "C", "", "Use this as current directory")
}

var rootCmd = &cobra.Command{
	Use:   "git-ws",
	Short: "A program to manage git worktrees",
	Long: `GitWS is a CLI tool that allows users to easily manage Git Worktrees`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
