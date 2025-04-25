package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	cli "github.com/urfave/cli/v3"
)

func matchedGroups(regEx, url string) (paramsMap map[string]string) {
    var compRegEx = regexp.MustCompile(regEx)
    match := compRegEx.FindStringSubmatch(url)

    paramsMap = make(map[string]string)
    for i, name := range compRegEx.SubexpNames() {
        if i > 0 && i <= len(match) {
            paramsMap[name] = match[i]
        }
    }
    return paramsMap
}

func cloneRepository(repository string, directory string, branch string) {
	if _, err := os.Stat(directory); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(directory, os.ModePerm); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Creating directory: %s\n", directory)
			ps := exec.Command("git", "clone", repository, directory, "--branch", branch)
			out, err := ps.CombinedOutput()
			if err != nil {
				log.Fatal(string(out))
			}
		} else {
			fmt.Printf("Already exists: %s\n", directory)
		}
	}
}

func remoteRepository(repository string) (string, string, string) {
	repo_groups := matchedGroups(`(?P<method>https?|git|ssh)(?::\/\/(?:\w+@)?|@)(?P<domain>.*?)(?:\.org|\.com)(?:\/|:)(?P<fullname>(?P<owner>.+?)\/(?P<repo>.+?))(?:\.git|\/)?$`, repository)
	if len(repo_groups) == 0 {
		err := fmt.Errorf("invalid repository: %s", repository)
		fmt.Println(err)
		os.Exit(1)
	}
	gitws_domain := strings.Trim(repo_groups["domain"], " \t")
	gitws_owner := strings.Trim(repo_groups["owner"], " \t")
	gitws_repo := strings.Trim(repo_groups["repo"], " \t")
	return gitws_domain, gitws_owner, gitws_repo
}

func defaultBranch(remoteUrl string) string {
	ps := exec.Command("git", "ls-remote", "--symref", remoteUrl, "HEAD")
	out, err := ps.CombinedOutput()
	if err != nil {
		log.Fatal(string(out))
	}
	branchGroups := matchedGroups(`refs\/heads\/(?P<branch>.*) *HEAD.*`, string(out))
	rootBranch := strings.Trim(branchGroups["branch"], " \t")
	return rootBranch
}

func CloneCmd() *cli.Command{
	cmd := &cli.Command{
		Name:  "clone",
		Usage: "Clone a repository and prepare it with GitWS structure.",
		Action: CloneFn,
		ArgsUsage: "<repository>",
	}
	return cmd
}

func CloneFn(ctx context.Context, cmd *cli.Command) error {
	if foundConfig {
		log.Fatal("Already inside a GitWS workspace")
	}

	remoteUrl := strings.Trim(cmd.Args().First(), " \t")
	rootDomain, rootOwner, rootRepo := remoteRepository(remoteUrl)
	rootBranch := defaultBranch(remoteUrl)

	// Clone the repository
	rootDir := filepath.Join(currentDir, rootRepo)
	gitDir := filepath.Join(rootDir, rootBranch)
	cloneRepository(remoteUrl, gitDir, rootBranch)

	// Finally save the configuration file
	gitwsConfig = GitwsConfig{
		RemoteUrl: remoteUrl,
		RootDomain: rootDomain,
		RootOwner: rootOwner,
		RootRepo: rootRepo,
		RootBranch: rootBranch,
		RootDir: rootDir,
		GitDir: gitDir,
	}
	writeGitwsConfig(gitwsConfig.RootDir, gitwsConfig)
	return nil
}
