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


func AddCmd() *cli.Command{
	cmd := &cli.Command{
		Name:  "add",
		Usage: "Add a branch from remote to the current gitws",
		Action: AddFn,
	}
	return cmd
}

func AddFn(ctx context.Context, cmd *cli.Command) error {
	branches := gitListAllBranches()
	wsList := Map(branches, func(a string) GitwsWorktree {
		return GitwsWorktree{
			Branch: a,
			Path: "",
			Sha: "",
		}
	})
	iterChoice := iterative(wsList)
	worktreeDir, err := GitWorktreeAdd(iterChoice.Branch)
	if err != nil {
		return err
	}
	postCmd := fmt.Sprintf(`cd %s`, worktreeDir)
	SetPostCmd(wsTempCmd, postCmd)
	return nil
}

func gitListAllBranches() []string {
	ps := exec.Command("git", "-C", gitwsConfig.GitDir, "branch", "--remotes", "--format", "%(refname)")
	out, err := ps.CombinedOutput()
	if err != nil {
		log.Fatal(string(out))
	}
	var branches []string
	for line := range strings.SplitSeq(string(out), "\n") {
		branches_groups := matchedGroups(`^refs/remotes/origin/(?P<branch>.*)$`, line)
		if len(branches_groups) > 0 {
			if !strings.Contains(branches_groups["branch"], "HEAD") {
				branches = append(branches, branches_groups["branch"])
			}
		}
	}
	return branches
}

func GitWorktreeAdd(branch string) (string, error) {
	worktreeDir := filepath.Join(gitwsConfig.RootDir, branch)
	ps := exec.Command("git", "-C", gitwsConfig.GitDir, "worktree", "add", "-B", branch, worktreeDir, gitwsConfig.RootBranch)
	out, err := ps.CombinedOutput()
	if err != nil {
		log.Fatal(string(out))
		return "", err
	}
	return worktreeDir, nil
}
