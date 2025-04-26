package cmd

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"

	cli "github.com/urfave/cli/v3"
)

type GitwsWorktree struct {
	Sha    string
	Path   string
	Branch string
}

func ListCmd() *cli.Command{
	cmd := &cli.Command{
		Name:  "list",
		Usage: "List all worktrees that currently exist.",
		Action: ListFn,
	}
	return cmd
}

func ListFn(ctx context.Context, cmd *cli.Command) error {
	if !foundConfig {
		log.Fatal("Gitws config file not found")
	}
	
	wsList := GetWsList()
	maxWidth := MaxWidth(wsList)

	for _, worktree := range GetWsList() {
		fmt.Printf(" % *s : %s\n", maxWidth, worktree.Branch, worktree.Path)
	}
	return nil
}

func GetWsList() []GitwsWorktree {
	ps := exec.Command("git", "-C", gitwsConfig.GitDir, "worktree", "list")
	out, err := ps.CombinedOutput()
	if err != nil {
		log.Fatal(string(out))
	}
	worktreeList := []GitwsWorktree{}
	for line := range strings.SplitSeq(string(out), "\n") {
		groups := matchedGroups(`^(?P<path>\/(?:[^\/ ]+\/?)*) +(?P<sha>[a-f0-9]+) \[(?P<branch>.+)\]$`, line)
		if len(groups) > 0 {
			worktreeList = append(worktreeList, GitwsWorktree{
				Sha: groups["sha"],
				Path: groups["path"],
				Branch: groups["branch"],
			})
		}
	}
	return worktreeList
}
