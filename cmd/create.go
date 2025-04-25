package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new worktree from a branch name",
		Run: func(cmd *cobra.Command, args []string) {
			branchName := strings.Trim(args[0], " \t")
			config := readGitwsConfig(gitwsDir)
			worktreeDir := filepath.Join(config.RootDir, branchName)

			ps := exec.Command("git", "-C", config.GitDir, "worktree", "add", "-B", branchName, worktreeDir, config.RootBranch)
			out, err := ps.CombinedOutput()
			if err != nil {
				log.Fatal(string(out))
			}
			fmt.Printf("%s\n", string(out))
		},
	}
	return createCmd
}
