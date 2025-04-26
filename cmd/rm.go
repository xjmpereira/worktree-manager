package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	cli "github.com/urfave/cli/v3"
)

func RmCmd() *cli.Command{
	cmd := &cli.Command{
		Name:  "rm",
		Usage: "Remove a worktree and clean it up",
		Action: RmFn,
	}
	return cmd
}

func RmFn(ctx context.Context, cmd *cli.Command) error {
	wsList := GetWsList()
	iterChoice := iterative(wsList)
	if iterChoice.Path == gitwsConfig.GitDir {
		return fmt.Errorf("removal of root branch not allowed")
	} else {
		gitWorktreeRemove(iterChoice.Path)
		gitWorktreePrune()
		cleanNestedDirectories(iterChoice.Path)
	}
	return nil
}

func gitWorktreeRemove(worktreeDir string) {
	ps := exec.Command("git", "-C", gitwsConfig.GitDir, "worktree", "remove", "--force", worktreeDir)
	out, err := ps.CombinedOutput()
	if err != nil {
		log.Fatal(string(out))
	}
}

func gitWorktreePrune() {
	ps := exec.Command("git", "-C", gitwsConfig.GitDir, "worktree", "prune")
	out, err := ps.CombinedOutput()
	if err != nil {
		log.Fatal(string(out))
	}
}

func cleanNestedDirectories(dir string) error {
	for len(dir) > 1 {
		dir = filepath.Dir(dir)
		insideWS := strings.Compare(dir, gitwsConfig.RootDir) > 0
		if insideWS && emptyDir(dir) {
			err := os.Remove(dir)
			if err != nil {
				return nil
			}
		}
	}
	return nil
}

func emptyDir(dir string) bool {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	return len(files) == 0
}
