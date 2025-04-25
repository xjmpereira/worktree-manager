package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all worktrees that currently exist.",
		Run: func(cmd *cobra.Command, args []string) {
			config := readGitwsConfig(gitwsDir)
			fmt.Printf("Using GITWS config: %s\n", gitwsDir)
			ps := exec.Command("git", "-C", config.GitDir, "worktree", "list")
			out, err := ps.CombinedOutput()
			if err != nil {
				log.Fatal(string(out))
			}
			fmt.Printf("%s\n", string(out))
		},
	}
	return listCmd
}
