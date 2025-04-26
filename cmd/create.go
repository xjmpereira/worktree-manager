package cmd

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	cli "github.com/urfave/cli/v3"
)

func CreateCmd() *cli.Command{
	cmd := &cli.Command{
		Name:  "create",
		Usage: "Create a new worktree from a branch name",
		Action: CreateFn,
	}
	return cmd
}

func CreateFn(ctx context.Context, cmd *cli.Command) error {
	if !foundConfig {
		log.Fatal("Gitws config file not found")
	}
	branchName := strings.Trim(cmd.Args().First(), " \t")
	worktreeDir := filepath.Join(gitwsConfig.RootDir, branchName)

	ps := exec.Command("git", "-C", gitwsConfig.GitDir, "worktree", "add", "-B", branchName, worktreeDir, gitwsConfig.RootBranch)
	out, err := ps.CombinedOutput()
	if err != nil {
		log.Fatal(string(out))
	}
	fmt.Printf("%s\n", string(out))
	return nil
}

